package model

type Transfer struct {
	TransferID string  `json:"transferID" xorm:"varchar(255) pk"`
	PlayerID   string  `json:"playerID"   xorm:"varchar(30) notnull"`
	Type       string  `json:"type"   xorm:"varchar(20) notnull"`
	BetID      string  `json:"betID"   xorm:"varchar(64) notnull"`
	GameID     string  `json:"gameID"   xorm:"varchar(64) notnull"`
	Amount     float64 `json:"amount" xorm:"float notnull" `
	Success    bool    `json:"success" xorm:"tinyint(1) notnull"`
	Created    int64   `xorm:"bigint"  `
	Updated    int64   `xorm:"bigint"  `
}

func GetTransferBy(playerID string) (m []Transfer, err error) {
	session := MyDB.NewSession()
	defer session.Close()
	err = session.Where("PlayerID=?", playerID).Desc("Created").Limit(10, 0).Find(&m)
	return
}

func AddTransfer(m Transfer) (err error) {
	session := MyDB.NewSession()
	defer session.Close()
	_, err = session.Insert(&m)
	return
}
