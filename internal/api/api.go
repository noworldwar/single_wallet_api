package api

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noworldwar/single_wallet_api/internal/model"
	"github.com/noworldwar/single_wallet_api/internal/pkg/utils"
	"github.com/rs/xid"
)

func Validate(c *gin.Context) {
	playerID := c.PostForm("playerID")
	var player_data model.Player
	var err error

	if playerID != "" && c.PostForm("token") == "" {
		player_data, err = model.GetPlayer(playerID)
		if err != nil {
			c.JSON(400, gin.H{"message": "player not found"})
			return
		}
		token := model.SetPlayerInfo(player_data)
		c.JSON(200, gin.H{"token": token})
		return
	}

	if c.PostForm("token") != "" {
		player_info := model.GetPlayerInfo(c.PostForm("token"))
		player_data, err = model.GetPlayer(player_info.PlayerID)
		if err != nil {
			c.JSON(400, gin.H{"message": "player not found"})
			return
		}
		c.JSON(200, gin.H{"playerID": player_data.PlayerID, "nickname": player_data.Nickname, "currency": player_data.Currency, "test": utils.IntToBool(player_data.Test), "time": player_data.Created})
		return
	} else if playerID == "" && c.PostForm("token") == "" {
		c.JSON(400, gin.H{"message": "Missing parameter"})
		return
	}

	// c.JSON(200, gin.H{"playerID": player_data.PlayerID, "nickname": player_data.Nickname, "currency": player_data.Currency, "test": utils.IntToBool(player_data.Test), "time": player_data.Created})
}

func GetBalance(c *gin.Context) {
	// token := c.PostForm("token")
	playerID := c.PostForm("playerID")
	if missing := utils.CheckPostFormData(c, "token", "playerID"); missing != "" {
		c.JSON(400, gin.H{"message": "Missing parameter: " + missing})
		return
	}

	player_data, err := model.GetPlayer(playerID)
	if err != nil {
		c.JSON(400, gin.H{"message": "player not found"})
	}
	// player_info := model.GetPlayerInfo(token)
	// fmt.Println(player_info)
	c.JSON(200, gin.H{"balance": player_data.Balance, "currency": player_data.Currency, "time": player_data.Created})
}

func Debit(c *gin.Context) {
	playerID := c.PostForm("playerID")
	amount := c.PostForm("amount")
	amount_float, _ := strconv.ParseFloat(amount, 64)
	currency := c.PostForm("currency")
	balance, err1 := model.UpdateBalance(playerID, -amount_float)
	if err1 != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
	}
	refID := time.Now().Format("20060102") + xid.New().String()
	transfer := model.Transfer{
		TransferID: refID,
		PlayerID:   playerID,
		Type:       "Debit",
		Amount:     amount_float,
		Success:    true,
		Created:    time.Now().Unix(),
		Updated:    time.Now().Unix(),
	}
	err := model.AddTransfer(transfer)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
	}
	c.JSON(200, gin.H{"balance": balance, "currency": currency, "time": time.Now().Unix(), "refID": refID})
}

func Credit(c *gin.Context) {
	playerID := c.PostForm("playerID")
	amount := c.PostForm("amount")
	amount_float, _ := strconv.ParseFloat(amount, 64)
	currency := c.PostForm("currency")
	balance, err1 := model.UpdateBalance(playerID, amount_float)
	if err1 != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
	}
	refID := time.Now().Format("20060102") + xid.New().String()
	transfer := model.Transfer{TransferID: refID,
		PlayerID: playerID,
		Type:     "Credit",
		Amount:   amount_float,
		Success:  true,
		Created:  time.Now().Unix(),
		Updated:  time.Now().Unix(),
	}
	err := model.AddTransfer(transfer)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
	}
	c.JSON(200, gin.H{"balance": balance, "currency": currency, "time": time.Now().Unix(), "refID": refID})
}
