package main

import (
	"github.com/noworldwar/myapi/app"
)

func main() {
	app.InitMySQL()
	app.InitRedis()
	app.InitRouter()
	go app.RunRouter()
	app.Cleanup()
}
