package app

import (
	"log"

	"github.com/noworldwar/single_wallet_api/internal/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func InitMySQL() {
	model.MyDB, _ = xorm.NewEngine("mysql", "root:password@tcp(127.0.0.1:3306)/mysite?charset=utf8mb4")
	model.MyDB.SetMapper(core.SameMapper{})

	err := model.MyDB.Ping()
	if err != nil {
		log.Fatalln("Init MySQL Error:", err)
	}

	AutoMigrate()
}

func AutoMigrate() {
	err := model.MyDB.Sync2(new(model.Player), new(model.Transfer), new(model.Wallet))
	if err != nil {
		log.Fatalln("User Sync Error:", err)
	}
}
