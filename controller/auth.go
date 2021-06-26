package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noworldwar/myapi/entity"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("wgtoken"); err == nil {
			auth, _ := entity.GetAuth(cookie.Value)
			c.Set("Auth", auth)
		}
		c.Next()
	}
}
