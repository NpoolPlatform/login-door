package redis

import (
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/go-redis/redis/v8"
)

func Client() *redis.Client {
	return app.Redis().Client
}

func InsertKeyInfo(keyword, userID string, info interface{}, ttl time.Duration)
