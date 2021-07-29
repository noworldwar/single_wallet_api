package model

import "time"

type Player struct {
	PlayerID   string    `json:"playerID" xorm:"varchar(30) pk"`
	OpPlayerID string    `json:"opPlayerID" xorm:"varchar(25) notnull"`
	Nickname   string    `json:"nickname" xorm:"varchar(30)"`
	Currency   string    `json:"currency" xorm:"varchar(5) notnull"`
	Password   string    `json:"password" xorm:"varchar(255)"`
	Balance    int64     `json:"balance"`
	Test       int       `xorm:"int notnull"`
	Disabled   bool      `json:"disabled"`
	Created    time.Time `json:"created"  xorm:"created"`
	Updated    time.Time `json:"updated"  xorm:"updated"`
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
