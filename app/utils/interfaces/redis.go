package interfaces

// Redis ...
type Redis interface {
	Set(key string, values string) error
	Get(key string) (string, error)
	HSet(key string, data map[string]string) error
	HGet(key string, field string) (string, error)
	HGetAll(key string) (map[string]string, error)
	SAdd(key string, member ...string) error
	SMembers(key string) ([]string, error)
}
