package settings

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Redis ...
type Redis struct {
	client  *redis.Client
	context context.Context
}

// Initialize ...
func (Redis Redis) Initialize() Redis {
	Redis.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	Redis.context = Redis.client.Context()

	return Redis
}

// Set ...
func (Redis Redis) Set(key string, value string) error {
	return Redis.client.Set(Redis.context, key, value, 0).Err()
}

// Get ...
func (Redis Redis) Get(key string) (string, error) {
	return Redis.client.Get(Redis.context, key).Result()
}

// HSet ...
func (Redis Redis) HSet(key string, data map[string]string) error {
	return Redis.client.HSet(Redis.context, key, data).Err()
}

// HGet ...
func (Redis Redis) HGet(key string, field string) (string, error) {
	return Redis.client.HGet(Redis.context, key, field).Result()
}

// HGetAll ...
func (Redis Redis) HGetAll(key string) (map[string]string, error) {
	return Redis.client.HGetAll(Redis.context, key).Result()
}

// SAdd ...
func (Redis Redis) SAdd(key string, member ...string) error {
	return Redis.client.SAdd(Redis.context, key, member).Err()
}

// SMembers ...
func (Redis Redis) SMembers(key string) ([]string, error) {
	return Redis.client.SMembers(Redis.context, key).Result()
}
