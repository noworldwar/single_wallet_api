package model

type Wallet struct {
	Balance float64 `json:"balance"`
	Success bool    `json:"success"`
}

func UpdateBalance(playerID string, amount float64) (float64, error) {
	session := MyDB.NewSession()
	defer session.Close()
	_, err := session.Exec("UPDATE Player SET Balance=Balance+? WHERE PlayerID=?", amount, playerID)
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
