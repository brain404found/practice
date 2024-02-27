package natsClient

import (
	"context"
	"encoding/json"
	"fmt"
	"service-1/tracer"

	"service-1/glog"

	"github.com/bankplanet9/goshared/v2/dto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/go-playground/validator/v10"
	"github.com/nats-io/nats.go"
)

// EventHandlerFunc is a type of pub-sub handlers
type EventHandlerFunc func(context.Context, *nats.Msg)

// HandlerFunc is a type of request-reply handlers
type HandlerFunc func(context.Context, *nats.Msg) *dto.Response

// WrapEventHandler выступает как middleware и оборачивает fn в функцию-обработчик NATS типа func(*nats.Msg),
// обрабатывая входящее сообщение и валидирую его перед тем, как вызвать fn
func WrapEventHandler[T interface{}](fn func(ctx context.Context, req T) error) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		trc := otel.Tracer("NATS Sub")
		carrier := tracer.NewNATSHeaderCarrier(&msg.Header)
		ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)
		ctx, span := trc.Start(ctx, fmt.Sprintf("[SUB] %v (%v)", msg.Subject, GetFunctionName(fn)),
			trace.WithAttributes(semconv.MessagingOperationReceive))
		defer span.End()

		log := glog.NewTracingLogger()

		var body T
		err := json.Unmarshal(msg.Data, &body)
		if err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("Cannot unmarshal incoming message")
			span.SetStatus(codes.Error, err.Error())
			msg.Term()
			return
		}

		validate := validator.New()
		err = validate.Struct(body)
		if err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("validation error on incoming message")
			span.SetStatus(codes.Error, err.Error())
			msg.Term()
			return
		}

		err = fn(ctx, body)
		if err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("error occurred while handler execution")
			span.SetStatus(codes.Error, err.Error())
			msg.Nak()
			return
		}
	}
}

// WrapHandler выступает как middleware и оборачивает fn в функцию-обработчик NATS типа func(*nats.Msg),
// обрабатывая входящее сообщение и валидирую его перед тем, как вызвать fn
//
// WrapHandler берет на себя отправку ответа на запрос с помощью NATS, пользуясь структурой dto.Response
// Возможны следующие ошибки в ответе:
//
// 400 - Не удалось десериализовать полученное сообщение
//
// 422 - Полученное сообщение не является валидным
//
// 500 - Невозможно сериализовать ответ fn в JSON
func WrapHandler[T interface{}](fn func(ctx context.Context, req T) *dto.Response) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		defer msg.Ack()

		var payload map[string]interface{}
		_ = json.Unmarshal(msg.Data, &payload)

		trc := otel.Tracer("NATS Rep")
		carrier := tracer.NewNATSHeaderCarrier(&msg.Header)
		ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)
		ctx, span := trc.Start(ctx, fmt.Sprintf("[REP] %v (%v)", msg.Subject, GetFunctionName(fn)),
			trace.WithAttributes(semconv.MessagingOperationReceive))
		defer span.End()

		log := glog.NewTracingLogger()

		var body T
		err := json.Unmarshal(msg.Data, &body)
		if err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("cannot unmarshal incoming message")
			span.SetStatus(codes.Error, err.Error())
			reply(msg, dto.NewErrorResponse(400, err.Error()))
			return
		}

		validate := validator.New()
		err = validate.Struct(body)
		if err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("validation error on incoming message")
			span.SetStatus(codes.Error, err.Error())
			reply(msg, dto.NewErrorResponse(422, err.Error()))
			return
		}

		response := fn(ctx, body)
		v, err := json.Marshal(response)
		if err != nil {
			log.Error().Ctx(ctx).Err(err).Msg("cannot convert response to json")
			span.SetStatus(codes.Error, err.Error())
			reply(msg, dto.NewErrorResponse(500, err.Error()))
			return
		}
		span.SetStatus(codes.Ok, "complete")
		reply(msg, v)
	}
}

type Client struct {
	ServiceName string
	Nc          *nats.Conn
	Js          nats.JetStreamContext
}

func NewNATSClient(serviceName, url string) *Client {
	nc, err := nats.Connect(url)
	if err != nil {
		panic("Cannot connect to NATS Server. Error: " + err.Error())
	}
	js, _ := nc.JetStream()
	if err != nil {
		panic("Cannot create JetStream connection. Error: " + err.Error())
	}

	return &Client{
		ServiceName: serviceName,
		Nc:          nc,
		Js:          js,
	}
}

func (c *Client) Publish(ctx context.Context, subject string, data interface{}, headers map[string]string) error {
	msg := nats.NewMsg(subject)

	// Adding headers
	for k, v := range headers {
		msg.Header.Add(k, v)
	}

	trc := otel.Tracer("NATS Pub")
	ctx, span := trc.Start(ctx, fmt.Sprintf("[PUB] %v", subject),
		trace.WithAttributes(semconv.MessagingOperationPublish))
	defer span.End()
	carrier := tracer.NewNATSHeaderCarrier(&msg.Header)
	otel.GetTextMapPropagator().Inject(ctx, carrier)

	jsonData, err := json.Marshal(data)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	msg.Data = jsonData

	_, err = c.Js.PublishMsg(msg)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}

func (c *Client) Subscribe(subject string, callback func(*nats.Msg)) error {
	_, err := c.Js.Subscribe(subject, callback, nats.Durable(c.ServiceName))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() {
	c.Nc.Close()
}

func (c *Client) Request(ctx context.Context, subject string, req interface{}, headers map[string]string) (resp []byte, err error) {
	msg := nats.NewMsg(subject)

	// Adding headers
	for k, v := range headers {
		msg.Header.Add(k, v)
	}

	trc := otel.Tracer("NATS Req")
	ctx, span := trc.Start(ctx, fmt.Sprintf("[REQ] %v", subject),
		trace.WithAttributes(semconv.MessagingOperationPublish))
	defer span.End()
	carrier := tracer.NewNATSHeaderCarrier(&msg.Header)
	otel.GetTextMapPropagator().Inject(ctx, carrier)

	// First, we convert dynamically-typed request body to map
	data, err := json.Marshal(&req)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	msg.Data = data
	span.SetAttributes(attribute.String("request", string(data)))

	respMsg, err := c.Nc.RequestMsgWithContext(ctx, msg)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	span.SetAttributes(attribute.String("response", string(respMsg.Data)))
	span.SetStatus(codes.Ok, "ok")
	return respMsg.Data, nil
}

// Shortcut function wrapper for nats.Msg.Respond
func reply(msg *nats.Msg, data []byte) error {
	err := msg.Respond(data)
	if err != nil {
		return err
	}
	return nil
}
