package app

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/noworldwar/myapi/controller"
	ctr "github.com/noworldwar/myapi/controller"
	"github.com/noworldwar/myapi/model"
)

func InitRouter() {
	r := gin.Default()
	r.Use(ctr.CheckToken())

	r.POST("/validate", controller.Validate)
	r.POST("/balance", controller.GetBalance)
	r.POST("/debit", controller.Debit)
	r.POST("/credit", controller.Credit)

	model.WGServer = http.Server{Addr: ":80", Handler: r}
}

func RunRouter() {
	if err := model.WGServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func loadTemplates() multitemplate.Renderer {

	r := multitemplate.NewRenderer()

	includes, err := filepath.Glob("view/page/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, include := range includes {
		r.AddFromFiles(filepath.Base(include), "view/layout/base.html", include)
	}

	return r
}
