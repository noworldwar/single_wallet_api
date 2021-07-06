package model

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var RDB *redis.Client

type PlayerInfo struct {
	PlayerID string
	Nickname string
}

func SetPlayerInfo(player Player) string {
	token := uuid.New().String()
	token = strings.Replace(token, "-", "", -1)
	b, _ := json.Marshal(PlayerInfo{PlayerID: player.PlayerID, Nickname: player.Nickname})
	_ = RDB.Set(context.Background(), token, string(b), time.Hour*1).Err()
	return token
}

func GetPlayerInfo(token string) (info PlayerInfo) {
	info = PlayerInfo{}
	val, err := RDB.Get(context.Background(), token).Result()
	if err != nil {
		return
	}
	err_1 := json.Unmarshal([]byte(val), &info)
	if err_1 != nil {
		return
	}
	_ = RDB.Expire(context.Background(), token, time.Hour*1).Err()
	return info
}
