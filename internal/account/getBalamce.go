package account

import (
	"database/sql"
	"fmt"
	"bank-account-system/db"
)

func GetBalance(accountID int) (float64, error) {
	var balance float64
	query := `SELECT balance FROM accounts WHERE id = $1`

	err := db.DB.QueryRow(query, accountID).Scan(&balance)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("account with id %d does not exist", accountID)
	}
	if err != nil {
		return 0, fmt.Errorf("could not fetch account balance: %v", err)
	}

	return balance, nil
}
