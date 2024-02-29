package tracer

import (
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/propagation"
)

// HeaderCarrier adapts http.Header to satisfy the otel.TextMapCarrier interface.
type NATSHeaderCarrier struct {
	nc *nats.Header
}

func NewNATSHeaderCarrier(headers *nats.Header) propagation.TextMapCarrier {
	return NATSHeaderCarrier{
		nc: headers,
	}
}

// Get returns the value associated with the passed key.
func (nc NATSHeaderCarrier) Get(key string) string {
	return nc.nc.Get(key)
}

// Set stores the key-value pair.
func (nc NATSHeaderCarrier) Set(key string, value string) {
	nc.nc.Set(key, value)
}

// Keys lists the keys stored in this carrier.
func (nc NATSHeaderCarrier) Keys() []string {
	keys := make([]string, 0, len(*nc.nc))
	for k := range *nc.nc {
		keys = append(keys, k)
	}
	return keys
}
