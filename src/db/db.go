package db

import (
	"log"
	"url-shortener/src/config"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func ConnectToRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr: config.ConfigVariables.RedisHost + ":" + config.ConfigVariables.RedisPort,
	})

	_, err := Client.Ping().Result()

	if err != nil {
		log.Fatalf("Could not connect to the Redis database: %v", err)
	}
}
