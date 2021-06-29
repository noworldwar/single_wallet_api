package api

import (
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	c.JSON(200, gin.H{"playerID": "test1", "nickname": "test1", "currency": "RMB", "test": false, "time": 1624944938})
	// playerID := c.PostForm("playerID")
	// player_data, err := model.GetPlayer(playerID)
	// if err != nil {
	// 	c.JSON(400, gin.H{"message": "player not found"})
	// }
	// c.JSON(200, gin.H{"playerID": player_data.PlayerID, "nickname": player_data.Nickname, "currency": player_data.Currency, "test": player_data.Test, "time": player_data.Created})
}
func GetBalance(c *gin.Context) {
	c.JSON(200, gin.H{"balacne": 3000000, "currency": "RMB", "time": 1624944938})
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
