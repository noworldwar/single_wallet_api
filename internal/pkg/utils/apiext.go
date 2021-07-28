package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, code int, msg string, err error) {
	errorMsg := msg
	if err != nil {
		errorMsg = fmt.Sprintf("%s: %v", msg, err)
	}
	WriteLog(msg)
	c.Set("ErrorMsg", errorMsg)
	c.JSON(code, gin.H{"error": msg})
}

func WriteLog(msg string) {
	f, err := os.OpenFile(fmt.Sprint(time.Now().Format("01-02-2006"))+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("\n" + msg); err != nil {
		panic(err)
	}

}
