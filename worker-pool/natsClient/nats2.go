package natsClient

import (
	"context"
	"fmt"
	"service-1/tracer"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

// EventHandlerFunc is a type of pub-sub handlers
type EventHandlerFuncV2 func(context.Context, *nats.Msg)

// WrapEventHandler выступает как middleware и оборачивает fn в функцию-обработчик NATS типа func(*nats.Msg),
// обрабатывая входящее сообщение и валидирую его перед тем, как вызвать fn
func WrapEventHandler_V2(fn EventHandlerFuncV2) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		trc := otel.Tracer("NATS Sub")
		carrier := tracer.NewNATSHeaderCarrier(&msg.Header)
		ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)
		ctx, span := trc.Start(ctx, fmt.Sprintf("[SUB] %v (%v)", msg.Subject, GetFunctionName(fn)),
			trace.WithAttributes(semconv.MessagingOperationReceive))
		defer span.End()
		fn(ctx, msg)
	}
}
