package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noworldwar/single_wallet_api/internal/model"
	"github.com/rs/xid"
)

func Validate(c *gin.Context) {

	playerID := c.PostForm("playerID")
	if playerID == "" && c.PostForm("token") != "" {
		playerID = c.PostForm("token")
	}

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

	c.JSON(200, gin.H{"playerID": player_data.PlayerID, "nickname": player_data.Nickname, "currency": player_data.Currency, "test": test_bool, "time": player_data.Created})
}
func GetBalance(c *gin.Context) {
	playerID := c.PostForm("playerID")
	player_data, err := model.GetPlayer(playerID)
	if err != nil {
		c.JSON(400, gin.H{"message": "player not found"})
	}
	c.JSON(200, gin.H{"balance": player_data.Balance, "currency": player_data.Currency, "time": player_data.Created})
}

func Debit(c *gin.Context) {
	fmt.Println("playerID: ", c.PostForm("playerID"))
	fmt.Println("token: ", c.PostForm("token"))
	fmt.Println("amount: ", c.PostForm("amount"))
	fmt.Println("currency: ", c.PostForm("currency"))
	playerID := c.PostForm("playerID")
	if playerID == "" && c.PostForm("token") != "" {
		playerID = c.PostForm("token")
	}
	amount := c.PostForm("amount")
	amount_float, _ := strconv.ParseFloat(amount, 64)
	currency := c.PostForm("currency")
	balance, err1 := model.UpdateBalance(playerID, -amount_float)
	fmt.Println(err1.Error())
	refID := time.Now().Format("20060102") + xid.New().String()
	transfer := model.Transfer{TransferID: refID, PlayerID: playerID, Type: "Debit", Amount: amount_float, Success: true, Created: time.Now().Unix(), Updated: time.Now().Unix()}
	err := model.AddTransfer(transfer)
	fmt.Println(err.Error())
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
	}
	c.JSON(200, gin.H{"balance": balance, "currency": currency, "time": time.Now().Unix(), "refID": refID})
}

func Credit(c *gin.Context) {
	playerID := c.PostForm("playerID")
	if playerID == "" && c.PostForm("token") != "" {
		playerID = c.PostForm("token")
	}
	amount := c.PostForm("amount")
	amount_float, _ := strconv.ParseFloat(amount, 64)
	currency := c.PostForm("currency")
	balance, err1 := model.UpdateBalance(playerID, amount_float)
	fmt.Println(err1.Error())
	refID := time.Now().Format("20060102") + xid.New().String()
	transfer := model.Transfer{TransferID: refID, PlayerID: playerID, Type: "Credit", Amount: amount_float, Success: true, Created: time.Now().Unix(), Updated: time.Now().Unix()}
	err := model.AddTransfer(transfer)
	fmt.Println(err.Error())
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
	}
	c.JSON(200, gin.H{"balance": balance, "currency": currency, "time": time.Now().Unix(), "refID": refID})
}
