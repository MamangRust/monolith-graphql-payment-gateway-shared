package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/MamangRust/monolith-graphql-payment-gateway-pkg/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// CacheStore represents a Redis cache store.
type CacheStore struct {
	redis  *redis.Client
	logger logger.LoggerInterface
}

// NewCacheStore creates a new instance of CacheStore using the given context,
// Redis client, and logger.
func NewCacheStore(redis *redis.Client, logger logger.LoggerInterface) *CacheStore {
	return &CacheStore{redis: redis, logger: logger}
}

// GetFromCache retrieves a cached item from Redis by the specified key.
// It returns a pointer to the item of type T and a boolean indicating whether the
// item was found in the cache. If the item is not found or an error occurs during
// retrieval or unmarshalling, it logs an error and returns nil and false.
func GetFromCache[T any](ctx context.Context, store *CacheStore, key string) (*T, bool) {
	cached, err := store.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, false
	}
	if err != nil {
		store.logger.Error("Redis get error", zap.Error(err), zap.String("cacheKey", key))
		return nil, false
	}

	var result T
	if err := json.Unmarshal([]byte(cached), &result); err != nil {
		store.logger.Error("Failed to unmarshal cache", zap.Error(err), zap.String("cacheKey", key))
		return nil, false
	}

	return &result, true
}

// SetToCache stores the given data in the Redis cache using the specified key and expiration duration.
// It accepts a CacheStore, a key string, a pointer to the data to be cached, and the expiration time as parameters.
// The data is marshaled into JSON format before being set in the cache.
// If marshaling fails, an error is logged, and the function returns without storing the data.
// If setting the data in the cache fails, an error is logged; otherwise, a debug log indicates successful caching.
func SetToCache[T any](ctx context.Context, store *CacheStore, key string, data *T, expiration time.Duration) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		store.logger.Error("Failed to marshal cache", zap.Error(err), zap.String("cacheKey", key))
		return
	}

	if err := store.redis.Set(ctx, key, jsonData, expiration).Err(); err != nil {
		store.logger.Error("Failed to set cache", zap.Error(err), zap.String("cacheKey", key))
	} else {
		store.logger.Debug("Successfully cached data",
			zap.String("cacheKey", key),
			zap.Duration("expiration", expiration))
	}
}

// DeleteFromCache removes the entry associated with the specified key from the Redis cache.
// It accepts a CacheStore and a key string as parameters.
// If the deletion process encounters an error, it logs the error message along with the cache key.
func DeleteFromCache(ctx context.Context, store *CacheStore, key string) {
	if err := store.redis.Del(ctx, key).Err(); err != nil {
		store.logger.Error("Failed to delete cache", zap.Error(err), zap.String("cacheKey", key))
	}
}
