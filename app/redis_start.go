package app

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/noworldwar/myapi/model"
)

func InitRedis() {
	model.RedisDB = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	_, err := model.RedisDB.Ping().Result()
	if err != nil {
		log.Fatalln("Init Redis Error:", err)
	}
}
