package natsClient

import (
	"context"

	"github.com/nats-io/nats.go"
)

type NatsClientInterface interface {
	Publish(ctx context.Context, subject string, data interface{}, headers map[string]string) error
	Subscribe(subject string, callback func(*nats.Msg)) error
	Request(ctx context.Context, subject string, req interface{}, headers map[string]string) (resp []byte, err error)
}
