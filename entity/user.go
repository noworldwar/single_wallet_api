package entity

import (
	"github.com/noworldwar/myapi/model"
)

func InsertUser(m model.User) error {
	_, err := model.MyDB.Insert(m)
	return err
}

func GetUser(UserID string) (*model.User, error) {
	m := new(model.User)
	_, err := model.MyDB.Where("UserID=?", UserID).Get(m)
	return m, err
}

func UpdateUser(m model.User) (int64, error) {
	affected, err := model.MyDB.Where("UserID=?", m.UserID).Update(m)
	return affected, err
}
