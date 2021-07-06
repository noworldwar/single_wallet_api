package app

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/noworldwar/single_wallet_api/internal/model"
)

func InitRedis() {
	model.RDB = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: 1})

	_, err := model.RDB.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("Init Redis Error:", err)
	}
}
