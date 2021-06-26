package controller

import "github.com/gin-gonic/gin"

func Validate(c *gin.Context) {
	c.JSON(200, gin.H{"response": 200})
}
func GetBalance(c *gin.Context) {
	c.JSON(200, gin.H{"response": 200})
}

func Debit(c *gin.Context) {
	c.JSON(200, gin.H{"response": 200})
}

func Credit(c *gin.Context) {
	c.JSON(200, gin.H{"response": 200})
}
