package kvstore


type KeyValueStore interface {
	Get(key string) (string, error)
	Set(key string, value string) error

	GetString(key string) string
	GetBool(key string) bool
	GetInt64(key string) int64
	GetFloat64(key string) float64
}