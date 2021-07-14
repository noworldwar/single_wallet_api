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

var whiteList []string = []string{
	"52.77.199.143",
	"13.251.118.6",
}

func InitRouter() {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/validate", checkWhiteList, api.Validate)
	r.POST("/balance", checkWhiteList, api.GetBalance)
	r.POST("/debit", checkWhiteList, api.Debit)
	r.POST("/credit", checkWhiteList, api.Credit)

	model.WGServer = http.Server{Addr: ":7901", Handler: r}
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

func checkWhiteList(c *gin.Context) {
	isLegalIp := false
	for _, v := range whiteList {
		if v == c.ClientIP() {
			isLegalIp = true
		}
	}
	if !isLegalIp {
		c.JSON(500, gin.H{"Message": "Permission Denied"})
		c.Abort()
	}
	c.Next()
}
