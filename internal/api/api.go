package api

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noworldwar/single_wallet_api/internal/model"
	"github.com/noworldwar/single_wallet_api/internal/pkg/utils"
	"github.com/rs/xid"
	logrus "github.com/sirupsen/logrus"
)

func Validate(c *gin.Context) {
	logrus.Println("token: ", c.PostForm("token"))
	logrus.Println("operatorID: ", c.PostForm("operatorID"))
	logrus.Println("appSecret: ", c.PostForm("appSecret"))
	playerID := c.PostForm("playerID")
	var player_data model.Player
	var err error

	// if playerID != "" && c.PostForm("token") == "" {
	// 	player_data, err = model.GetPlayer(playerID)
	// 	if err != nil {
	// 		c.JSON(400, gin.H{"message": "player not found"})
	// 		return
	// 	}

	// 	if player_data.PlayerID == "" {
	// 		c.JSON(400, gin.H{"message": "player not found"})
	// 		return
	// 	}
	// 	token := model.SetPlayerInfo(player_data)
	// 	c.JSON(200, gin.H{"token": token})
	// 	return
	// }

	// if c.PostForm("token") != "" {
	// 	player_info := model.GetPlayerInfo(c.PostForm("token"))
	// 	player_data, err = model.GetPlayer(player_info.PlayerID)
	// 	if err != nil || player_data.PlayerID == "" {
	// 		c.JSON(400, gin.H{"message": "player not found"})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{"playerID": player_data.PlayerID, "nickname": player_data.Nickname, "currency": player_data.Currency, "test": utils.IntToBool(player_data.Test), "time": player_data.Created})
	// 	return
	// } else if playerID == "" && c.PostForm("token") == "" {
	// 	c.JSON(400, gin.H{"message": "Missing parameter"})
	// 	return
	// }

	if playerID == "" && c.PostForm("token") != "" {
		playerID = c.PostForm("token")
	}
	player_data, err = model.GetPlayer(playerID)

	if err != nil || player_data.PlayerID == "" {
		c.JSON(400, gin.H{"message": "player not found"})
		return
	}

	c.JSON(200, gin.H{"playerID": player_data.PlayerID, "nickname": player_data.Nickname, "currency": player_data.Currency, "test": utils.IntToBool(player_data.Test), "time": player_data.Created})

	// c.JSON(200, gin.H{"playerID": player_data.PlayerID, "nickname": player_data.Nickname, "currency": player_data.Currency, "test": utils.IntToBool(player_data.Test), "time": player_data.Created})
}

func GetBalance(c *gin.Context) {
	logrus.Println("token: ", c.PostForm("token"))
	logrus.Println("operatorID: ", c.PostForm("operatorID"))
	logrus.Println("appSecret: ", c.PostForm("appSecret"))
	logrus.Println("playerID: ", c.PostForm("playerID"))
	// token := c.PostForm("token")

	playerID := c.PostForm("playerID")
	if playerID == "" && c.PostForm("token") != "" {
		playerID = c.PostForm("token")
	}

	player_data, err := model.GetPlayer(playerID)
	if err != nil || player_data.PlayerID == "" {
		c.JSON(400, gin.H{"message": "player not found"})
	}
	// player_info := model.GetPlayerInfo(token)
	// logrus.Println(player_info)
	c.JSON(200, gin.H{"balance": player_data.Balance, "currency": player_data.Currency, "time": player_data.Created})
}

func Debit(c *gin.Context) {
	logrus.Println("amount: ", c.PostForm("amount"))
	logrus.Println("playerID: ", c.PostForm("playerID"))
	logrus.Println("currency: ", c.PostForm("currency"))
	logrus.Println("token: ", c.PostForm("token"))
	logrus.Println("operatorID: ", c.PostForm("operatorID"))
	logrus.Println("appSecret: ", c.PostForm("appSecret"))
	logrus.Println("gameID: ", c.PostForm("gameID"))
	logrus.Println("betID: ", c.PostForm("betID"))
	logrus.Println("amount: ", c.PostForm("amount"))
	logrus.Println("currency: ", c.PostForm("currency"))
	logrus.Println("type: ", c.PostForm("type"))
	logrus.Println("time: ", c.PostForm("time"))

	playerID := c.PostForm("playerID")
	if playerID == "" && c.PostForm("token") != "" {
		playerID = c.PostForm("token")
	}
	amount := c.PostForm("amount")
	amount_float, _ := strconv.ParseFloat(amount, 64)
	currency := c.PostForm("currency")

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
	balance, _ := model.UpdateBalance(playerID, -amount_float)

	c.JSON(200, gin.H{"balance": balance, "currency": currency, "time": time.Now().Unix(), "refID": refID})
}

func Credit(c *gin.Context) {
	logrus.Println("token: ", c.PostForm("token"))
	logrus.Println("operatorID: ", c.PostForm("operatorID"))
	logrus.Println("appSecret: ", c.PostForm("appSecret"))
	logrus.Println("gameID: ", c.PostForm("gameID"))
	logrus.Println("uid: ", c.PostForm("uid"))
	logrus.Println("amount: ", c.PostForm("amount"))
	logrus.Println("uid: ", c.PostForm("uid"))
	logrus.Println("type: ", c.PostForm("type"))
	logrus.Println("time: ", c.PostForm("time"))
	logrus.Println("playerID: ", c.PostForm("playerID"))
	logrus.Println("currency: ", c.PostForm("currency"))

	playerID := c.PostForm("playerID")
	if playerID == "" && c.PostForm("token") != "" {
		playerID = c.PostForm("token")
	}
	amount := c.PostForm("amount")
	amount_float, _ := strconv.ParseFloat(amount, 64)
	currency := c.PostForm("currency")

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
	balance, _ := model.UpdateBalance(playerID, amount_float)

	c.JSON(200, gin.H{"balance": balance, "currency": currency, "time": time.Now().Unix(), "refID": refID})
}
