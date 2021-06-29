package model

import "time"

type Transfer struct {
	TransferID string    `json:"transferID" xorm:"varchar(30) pk"`
	PlayerID   string    `json:"playerID"   xorm:"varchar(30)"`
	Amount     int64     `json:"amount"`
	Success    bool      `json:"success"`
	Created    time.Time `json:"created"    xorm:"created"`
	Updated    time.Time `json:"updated"    xorm:"updated"`
}

func GetTransferBy(playerID string) (m []Transfer, err error) {
	session := MyDB.NewSession()
	defer session.Close()
	err = session.Where("player_id=?", playerID).Desc("created").Limit(10, 0).Find(&m)
	return
}

func AddTransfer(m Transfer) (err error) {
	session := MyDB.NewSession()
	defer session.Close()
	_, err = session.Insert(&m)
	return
}