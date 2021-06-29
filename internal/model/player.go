package model

type Player struct {
	PlayerID   string  `json:"playerID" xorm:"varchar(20)    pk"`
	OpPlayerID string  `json:"opPlayerID" xorm:"varchar(25) notnull"`
	Nickname   string  `json:"nickname" xorm:"varchar(255) notnull"`
	Currency   string  `json:"currency" xorm:"varchar(5) notnull"`
	Balance    float64 `json:"balance" xorm:"bigint         notnull"`
	Test       int     `xorm:"int notnull"`
	Created    int64   `xorm:"bigint"  `
	Updated    int64   `xorm:"bigint"  `
}

func GetPlayer(playerID string) (m Player, err error) {
	session := MyDB.NewSession()
	defer session.Close()
	_, err = session.ID(playerID).Get(&m)
	return
}

func AddPlayer(m Player) (err error) {
	session := MyDB.NewSession()
	defer session.Close()
	_, err = session.Insert(&m)
	return
}

func UpdatePlayer(m Player) (err error) {
	session := MyDB.NewSession()
	defer session.Close()
	_, err = session.ID(m.PlayerID).Update(&m)
	return
}
