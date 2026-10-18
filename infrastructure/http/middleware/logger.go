package middleware

import (
	"Momentum/infrastructure/http/request"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"go.opentelemetry.io/otel/trace"
)

func configureLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format: `{"time":"${time}","status":${status},"latency":"${latency}","method":"${method}","path":"${path}","request_id":"${request_id}","trace_id":"${trace_id}","span_id":"${span_id}","error":"${error}"}` + "\n",
		CustomTags: map[string]logger.LogFunc{
			"request_id": func(output logger.Buffer, ctx fiber.Ctx, _ *logger.Data, _ string) (int, error) {
				return output.WriteString(request.FromContext(ctx))
			},
			"trace_id": func(output logger.Buffer, ctx fiber.Ctx, _ *logger.Data, _ string) (int, error) {
				return output.WriteString(trace.SpanContextFromContext(ctx.Context()).TraceID().String())
			},
			"span_id": func(output logger.Buffer, ctx fiber.Ctx, _ *logger.Data, _ string) (int, error) {
				return output.WriteString(trace.SpanContextFromContext(ctx.Context()).SpanID().String())
			},
		},
	}))
}
