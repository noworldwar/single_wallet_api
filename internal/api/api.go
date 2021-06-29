package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/noworldwar/single_wallet_api/internal/model"
)

func Validate(c *gin.Context) {
	playerID := c.PostForm("playerID")
	player_data, err := model.GetPlayer(playerID)
	if err != nil {
		c.JSON(400, gin.H{"message": "player not found"})
	}
	var test_bool bool
	if player_data.Test == 0 {
		test_bool = false
	} else {
		test_bool = true
	}
	rsp := gin.H{"playerID": player_data.PlayerID, "nickname": player_data.Nickname, "currency": player_data.Currency, "test": test_bool, "time": player_data.Created}

	fmt.Println(rsp)
	c.JSON(200, rsp)
}
func GetBalance(c *gin.Context) {
	playerID := c.PostForm("playerID")
	player_data, err := model.GetPlayer(playerID)
	if err != nil {
		c.JSON(400, gin.H{"message": "player not found"})
	}
	rsp := gin.H{"balacne": player_data.Balance, "currency": player_data.Currency, "time": player_data.Created}
	fmt.Println(rsp)
	c.JSON(200, rsp)
}

func Debit(c *gin.Context) {
	c.JSON(200, gin.H{"balacne": 2000, "currency": "RMB", "time": 1574476825, "refID": "20200420XDCFSEDSE"})
}

func Credit(c *gin.Context) {
	c.JSON(200, gin.H{"balacne": 2000, "currency": "RMB", "time": 1574476825, "refID": "20200420XDCFSEDSE"})
}
