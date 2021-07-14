package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, code int, msg string, err error) {
	errorMsg := msg
	if err != nil {
		errorMsg = fmt.Sprintf("%s: %v", msg, err)
	}
	c.Set("ErrorMsg", errorMsg)
	c.JSON(code, gin.H{"error": msg})
}
