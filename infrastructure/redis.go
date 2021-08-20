package infrastructure

import "github.com/go-redis/redis/v7"

var RedisClient *redis.Client

func GetRedisClient() *redis.Client {
	return RedisClient
}

func SetRedisClient(client *redis.Client) {
	RedisClient = client
}
