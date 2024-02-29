package glog

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)


// TracingHook helps us to identify logs and find matching spans by their unique id
type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	spanId := getSpanIdFromContext(ctx) // as per your tracing framework
	e.Str("span-id", spanId)
}

func getSpanIdFromContext(ctx context.Context) string {
	span := trace.SpanFromContext(ctx)
	return span.SpanContext().SpanID().String()
}

func NewTracingLogger() zerolog.Logger {
	logger := zerolog.New(os.Stdout)
	logger = logger.Hook(TracingHook{})
	return logger
}

func NewLogger() zerolog.Logger {
	logger := zerolog.New(os.Stdout)
	return logger
}