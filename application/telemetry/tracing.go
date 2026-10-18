package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const tracerName = "Momentum/application"

func Start(ctx context.Context, name string) (context.Context, trace.Span) {
	return otel.Tracer(tracerName).Start(ctx, name)
}
