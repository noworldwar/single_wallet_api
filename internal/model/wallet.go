package model

type Wallet struct {
	Bank    string `json:"bank"`
	Balance int64  `json:"balance"`
	Success bool   `json:"success"`
}

func UpdateBalance(playerID string, amount int64) (int64, error) {
	session := MyDB.NewSession()
	defer session.Close()
	_, err := session.Exec("UPDATE player SET balance=balance+? WHERE player_id=?", amount, playerID)
	if err != nil {
		return 0, err
	}

	var m Player
	_, err = session.ID(playerID).Get(&m)
	if err != nil {
		return 0, err
	}

	return m.Balance, nil
}
