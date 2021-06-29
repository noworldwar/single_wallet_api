package model

import "errors"

type Wallet struct {
	Balance float64 `json:"balance"`
	Success bool    `json:"success"`
}

func UpdateBalance(playerID string, amount float64) (float64, error) {
	session := MyDB.NewSession()
	defer session.Close()
	affected, err := session.Exec("UPDATE Player SET Balance=Balance+? WHERE PlayerID=?", amount, playerID)
	if err != nil {
		return 0, err
	}

	if row, _ := affected.RowsAffected(); row == 0 {
		return 0, errors.New("affected 0")
	}

	var m Player
	_, err = session.ID(playerID).Get(&m)
	if err != nil {
		return 0, err
	}

	return m.Balance, nil
}
