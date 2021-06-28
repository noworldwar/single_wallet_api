package main

import (
	"single_wallet_api/internal/app"
)

func main() {
	app.InitMySQL()
	app.InitRedis()
	app.InitRouter()
	go app.RunRouter()
	app.Cleanup()
}
