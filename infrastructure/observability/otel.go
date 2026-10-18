package observability

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

const instrumentationName = "Momentum"

type Providers struct {
	tracerProvider *trace.TracerProvider
	meterProvider  *metric.MeterProvider
}

func Setup(ctx context.Context, serviceName, environment, endpoint string) (*Providers, error) {
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	if endpoint == "" {
		return &Providers{}, nil
	}

	serviceResource := resource.NewWithAttributes(
		"",
		attribute.String("service.name", serviceName),
		attribute.String("deployment.environment.name", environment),
	)
	res, err := resource.Merge(resource.Default(), serviceResource)
	if err != nil {
		return nil, fmt.Errorf("create telemetry resource: %w", err)
	}

	traceExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpointURL(endpoint),
	)
	if err != nil {
		return nil, fmt.Errorf("create OTLP trace exporter: %w", err)
	}

	metricExporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpointURL(endpoint),
	)
	if err != nil {
		_ = traceExporter.Shutdown(ctx)
		return nil, fmt.Errorf("create OTLP metric exporter: %w", err)
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(res),
	)
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
		metric.WithResource(res),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetMeterProvider(meterProvider)

	return &Providers{
		tracerProvider: tracerProvider,
		meterProvider:  meterProvider,
	}, nil
}

func (providers *Providers) Shutdown(ctx context.Context) error {
	if providers == nil {
		return nil
	}

	var shutdownErrors []error
	if providers.meterProvider != nil {
		shutdownErrors = append(shutdownErrors, providers.meterProvider.Shutdown(ctx))
	}
	if providers.tracerProvider != nil {
		shutdownErrors = append(shutdownErrors, providers.tracerProvider.Shutdown(ctx))
	}

	return errors.Join(shutdownErrors...)
}
