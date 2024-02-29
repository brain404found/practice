package kvstore

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nats-io/nats.go"
)

type natsKvStore struct {
	conn *nats.Conn
	js nats.JetStreamContext
	kv nats.KeyValue
}


func NewNATSKVStoreFromConn(conn *nats.Conn) KeyValueStore {
	js, err := conn.JetStream()
	if err != nil {
		panic("Cannot get JetStream context within NATS connection. Error: %v" + err.Error())
	}

	kv, err := js.KeyValue("sharedConfig")
	if err != nil {
		panic("Cannot get sharedConfig kv bucket. Error: " + err.Error())
	}

	return &natsKvStore{
		conn: conn,
		js: js,
		kv: kv,
	}
}

func NewNATSKVStore(natsUrl string) KeyValueStore {
	conn, err := nats.Connect(natsUrl)
	if err != nil {
		panic("Cannot connect to NATS KeyValue store. Error: %v" + err.Error())
	}
	js, err := conn.JetStream()
	if err != nil {
		panic("Cannot get JetStream context within NATS connection. Error: %v" + err.Error())
	}
	kv, err := js.KeyValue("sharedConfig")
	if err != nil {
		panic("Cannot get sharedConfig kv bucket. Error: " + err.Error())
	}
	return &natsKvStore{
		conn: conn,	
		js: js,
		kv: kv,
	}
}


func (s *natsKvStore) Get(key string) (string, error) {
	entry, err := s.kv.Get(key)
	if err != nil {
		return "", err
	}
	return string(entry.Value()), nil
}


func (s *natsKvStore) Set(key string, value string) error {
	_, err := s.kv.PutString(key, value)
	return err
}

// GetString returns the string value for the key if provided.
// 
// Panics if the key is not present
func (s *natsKvStore) GetString(key string) string {
	entry, err := s.kv.Get(key)
	if err != nil {
		panic("Cannot get value for key '" + key + "'. Error: " + err.Error())
	}
	return string(entry.Value())
}


// GetBool returns the boolean value for the key if provided.
// '1', 'true' or 'True' stored in KV storage will be returned as boolean value true.
// Any other value will be parsed as false.
// 
// Panics if the key is not present
func (s *natsKvStore) GetBool(key string) bool {
	entry, err := s.kv.Get(key)
	if err != nil {
		panic("Cannot get value for key '" + key + "'. Error: " + err.Error())
	}

	val := string(entry.Value())
	if val == "1" || strings.ToLower(val) == "true" {
		return true
	}
	return false
}


// GetInt64 returns the int64 value for the key if provided.
// Tries to parse int64 value from the key-value pair stored
// 
// Panics if the key is not present or if the value cannot be converted to int64
func (s *natsKvStore) GetInt64(key string) int64 {
	entry, err := s.kv.Get(key)
	if err != nil {
		panic("Cannot get value for key '" + key + "'. Error: " + err.Error())
	}
	
	valString := string(entry.Value())
	val, err := strconv.Atoi(valString)
	if err != nil {
		panic(fmt.Sprintf("Cannot convert KV %v=%v to int64", key, valString))
	}
	return int64(val)
}


// GetFloat64 returns the float64 value for the key if provided.
// Tries to parse float64 value from the key-value pair stored
// 
// Panics if the key is not present or if the value cannot be converted to float64
func (s *natsKvStore) GetFloat64(key string) float64 {
	entry, err := s.kv.Get(key)
	if err != nil {
		panic("Cannot get value for key '" + key + "'. Error: " + err.Error())
	}
	
	valString := string(entry.Value())
	val, err := strconv.ParseFloat(valString, 64)
	if err != nil {
		panic(fmt.Sprintf("Cannot convert KV %v=%v to float64", key, valString))
	}
	return float64(val)
}