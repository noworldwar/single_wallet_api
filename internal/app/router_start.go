package app

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/noworldwar/single_wallet_api/internal/api"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/noworldwar/single_wallet_api/internal/model"
)

func InitRouter() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/validate", api.Validate)
	r.POST("/balance", api.GetBalance)
	r.POST("/debit", api.Debit)
	r.POST("/credit", api.Credit)

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
