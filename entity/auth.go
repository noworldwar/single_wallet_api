package entity

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/noworldwar/myapi/model"
)

func SetAuth(token string, in *model.User) error {
	m := &model.Auth{
		UserID: in.UserID,
		Token:  token,
	}

	out, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return model.RedisDB.Set(token, string(out), 1*time.Hour).Err()
}

func ReSetAuth(m *model.Auth) error {
	out, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return model.RedisDB.Set(m.Token, string(out), 1*time.Hour).Err()
}

func GetAuth(token string) (string, error) {
	val, err := model.RedisDB.Get(token).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func ToAuth(val interface{}) *model.Auth {

	if val == nil {
		return nil
	}

	val_str := fmt.Sprintf("%s", val)

	out := &model.Auth{}
	err := json.Unmarshal([]byte(val_str), out)
	if err != nil {
		fmt.Println("Parse UserAuth Error:", err)
		return nil
	}
	return out
}
