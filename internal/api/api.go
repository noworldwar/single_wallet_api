package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	fmt.Println("playerID: ", c.PostForm("playerID"))
	fmt.Println("token: ", c.PostForm("token"))
	fmt.Println("operatorID: ", c.PostForm("operatorID"))
	fmt.Println("appSecret: ", c.PostForm("appSecret"))
	c.JSON(200, gin.H{"playerID": "mike", "nickname": "Mike", "currency": "RMB", "test": false, "time": 1624944938})
	// playerID := c.PostForm("playerID")
	// player_data, err := model.GetPlayer(playerID)
	// if err != nil {
	// 	c.JSON(400, gin.H{"message": "player not found"})
	// }
	// var test_bool bool
	// if player_data.Test == 0 {
	// 	test_bool = false
	// } else {
	// 	test_bool = true
	// }

	// c.JSON(200, gin.H{"playerID": player_data.PlayerID, "nickname": player_data.Nickname, "currency": player_data.Currency, "test": test_bool, "time": player_data.Created})
}
func GetBalance(c *gin.Context) {
	fmt.Println("playerID: ", c.PostForm("playerID"))
	fmt.Println("token: ", c.PostForm("token"))
	fmt.Println("operatorID: ", c.PostForm("operatorID"))
	fmt.Println("appSecret: ", c.PostForm("appSecret"))
	c.JSON(200, gin.H{"balacne": 30000000, "currency": "RMB", "time": 1624944938})
	// playerID := c.PostForm("playerID")
	// player_data, err := model.GetPlayer(playerID)
	// if err != nil {
	// 	c.JSON(400, gin.H{"message": "player not found"})
	// }
	// c.JSON(200, gin.H{"balacne": player_data.Balance, "currency": player_data.Currency, "time": player_data.Created})
}

func Debit(c *gin.Context) {
	c.JSON(200, gin.H{"balacne": 2000, "currency": "RMB", "time": 1574476825, "refID": "20200420XDCFSEDSE"})
}

func Credit(c *gin.Context) {
	c.JSON(200, gin.H{"balacne": 2000, "currency": "RMB", "time": 1574476825, "refID": "20200420XDCFSEDSE"})
}
