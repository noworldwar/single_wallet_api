package model

import "time"

type Player struct {
	PlayerID   string    `xorm:"varchar(20)    pk"`
	OpPlayerID string    `xorm:"varchar(25) notnull"`
	Nickname   string    `xorm:"varchar(255) notnull"`
	Currency   string    `xorm:"varchar(5) notnull"`
	Balance    int64     `xorm:"bigint         notnull"`
	Test       int       `xorm:"int notnull"`
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
