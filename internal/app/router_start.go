package app

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/noworldwar/single_wallet_api/internal/api"

	"github.com/gin-gonic/gin"
	"github.com/noworldwar/single_wallet_api/internal/model"
)

var whiteList []string = []string{
	"52.77.199.143",
	"13.251.118.6",
	"::1",
}

func InitRouter() {
	r := gin.Default()

	r.Use(cors.Default())
	r.Use(Logger())
	r.Use(gin.Recovery())

	group := r.Group("/api")

	group.POST("/validate", checkWhiteList, api.Validate)
	group.POST("/balance", checkWhiteList, api.GetBalance)
	group.POST("/debit", checkWhiteList, api.Debit)
	group.POST("/credit", checkWhiteList, api.Credit)
	group.POST("/rollback", checkWhiteList, api.Rollback)

	model.WGServer = http.Server{Addr: ":7901", Handler: r}
}

func RunRouter() {
	if err := model.WGServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func checkWhiteList(c *gin.Context) {
	// isLegalIp := false
	// for _, v := range whiteList {
	// 	if v == c.ClientIP() {
	// 		isLegalIp = true
	// 	}
	// }
	// if !isLegalIp {
	// 	c.JSON(500, gin.H{"Message": "Permission Denied"})
	// 	c.Abort()
	// }
	c.Next()
}
