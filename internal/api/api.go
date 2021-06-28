package api

import "github.com/gin-gonic/gin"

func Validate(c *gin.Context) {
	c.JSON(200, gin.H{"playerID": "C2R0289KK616RUHLR110", "nickname": "test1", "currency": "RMB", "test": false, "time": 1574476825})
}
func GetBalance(c *gin.Context) {
	c.JSON(200, gin.H{"balacne": 2000, "currency": "RMB", "time": 1574476825})
}

func Debit(c *gin.Context) {
	c.JSON(200, gin.H{"balacne": 2000, "currency": "RMB", "time": 1574476825, "refID": "20200420XDCFSEDSE"})
}

func Credit(c *gin.Context) {
	c.JSON(200, gin.H{"balacne": 2000, "currency": "RMB", "time": 1574476825, "refID": "20200420XDCFSEDSE"})
}
