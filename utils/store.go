package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	log.Println("redis 서버 연결중: ", os.Getenv("REDIS_HOST"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return rdb
}

func SetKey(ctx *context.Context, rdb *redis.Client, key string, value string, ttl int) {
	log.Println("REDIS key: ", key, ", value: ", value)
	rdb.Set(*ctx, key, value, 0)
	log.Println("키 설정 완료")
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
