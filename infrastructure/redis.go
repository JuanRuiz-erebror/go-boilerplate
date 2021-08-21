package infrastructure

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

var RedisClient *redis.Client

//var redisClient *redis.Client

func MainRedis() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "redis:6379"
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
		DB:   0,
	})
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(pong, err)

	fmt.Printf("start redis")

	//SetRedisClient(redisClient)
}

// func GetRedisClient() *redis.Client {
// 	return RedisClient
// }

// func SetRedisClient(client *redis.Client) {
// 	RedisClient = client
// }
