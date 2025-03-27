package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	log.Println("redis 서버 연결중: ", ReadEnv("REDIS_HOST"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     ReadEnv("REDIS_HOST"),
		Password: ReadEnv("REDIS_PASSWORD"),
		DB:       0,
	})

	return rdb
}

func SetKey(ctx *context.Context, rdb *redis.Client, key string, value string, ttl int) {
	rdb.Set(*ctx, key, value, 0)
}

func GetOriginURL(ctx *context.Context, rdb *redis.Client, shortURL string) (string, error) {
	originUrl, err := rdb.Get(*ctx, shortURL).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("원본 URL을 찾을 수 없습니다.")
	} else if err != nil {
		return "", fmt.Errorf("Redis 검색 실패: %v", err)
	}
	return originUrl, nil
}
