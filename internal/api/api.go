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

	var player_data model.Player
	var err error

	player_info := model.GetPlayerInfo(c.PostForm("token"))
	player_data, err = model.GetPlayer(player_info.PlayerID)
	if err != nil || player_data.PlayerID == "" {
		utils.ErrorResponse(c, 400, "Player not found: ", err)
		return
	}
	c.JSON(200, gin.H{"playerID": player_data.PlayerID, "nickname": player_data.Nickname, "currency": player_data.Currency, "test": utils.IntToBool(player_data.Test), "time": player_data.Created})
}

func GetBalance(c *gin.Context) {

	player_info := model.GetPlayerInfo(c.PostForm("token"))

	player_data, err := model.GetPlayer(player_info.PlayerID)
	if err != nil || player_data.PlayerID == "" {
		utils.ErrorResponse(c, 400, "Player not found: ", err)
		return
	}
	// player_info := model.GetPlayerInfo(token)
	// logrus.Println(player_info)
	c.JSON(200, gin.H{"balance": player_data.Balance, "currency": player_data.Currency, "time": player_data.Created})
}

func Debit(c *gin.Context) {

	if c.PostForm("token") == "" {
		c.JSON(404, gin.H{"message": "Token has expired"})
		return
	}

	playerID := c.PostForm("playerID")
	amount := c.PostForm("amount")
	amount_int, _ := strconv.ParseInt(amount, 10, 64)
	currency := c.PostForm("currency")

	isExist, err := model.CheckIfTransferExist(c.PostForm("betID"))
	if isExist {
		utils.ErrorResponse(c, 409, "Duplicate transaction: ", err)
		return
	}
	refID := time.Now().Format("20060102") + xid.New().String()
	transfer := model.BetTransfer{
		TransferID: refID,
		PlayerID:   playerID,
		Type:       "Debit",
		BetID:      c.PostForm("betID"),
		GameID:     c.PostForm("gameID"),
		Amount:     amount_int,
		Success:    true,
		Created:    time.Now(),
		Updated:    time.Now(),
	}
	AddErr := model.AddTransfer(transfer)
	if AddErr != nil {
		utils.ErrorResponse(c, 500, "Internal Server Error: ", AddErr)
		return
	}
	balance, _ := model.UpdateBalance(playerID, -amount_int)

	c.JSON(200, gin.H{"balance": balance, "currency": currency, "time": time.Now().Unix(), "refID": refID})
}

func Credit(c *gin.Context) {

	playerID := c.PostForm("playerID")
	amount := c.PostForm("amount")
	amount_int, _ := strconv.ParseInt(amount, 10, 64)
	currency := c.PostForm("currency")

	refID := time.Now().Format("20060102") + xid.New().String()
	transfer := model.BetTransfer{TransferID: refID,
		PlayerID: playerID,
		Type:     "Credit",
		BetID:    c.PostForm("betID"),
		GameID:   c.PostForm("gameID"),
		Amount:   amount_int,
		Success:  true,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	err := model.AddTransfer(transfer)
	if err != nil {
		utils.ErrorResponse(c, 500, "Internal Server Error: ", err)
		return
	}
	balance, _ := model.UpdateBalance(playerID, amount_int)

	c.JSON(200, gin.H{"balance": balance, "currency": currency, "time": time.Now().Unix(), "refID": refID})
}
