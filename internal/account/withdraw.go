package account

import (
	"fmt"

	"bank-account-system/db"

)

func Withdraw(accountID int, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("the amount must be positive")
	}

	var exists bool
	err := db.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM accounts WHERE id = $1)`, accountID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("could not check account existence: %v", err)
	}
	if !exists {
		return fmt.Errorf("account with id %d does not exist", accountID)
	}

	var balance float64
	query := `SELECT balance FROM accounts WHERE id = $1`
	err = db.DB.QueryRow(query, accountID).Scan(&balance)
	if err != nil {
		return fmt.Errorf("could not fetch account balance: %v", err)
	}

	if balance < amount {
		return fmt.Errorf("insufficient funds")
	}

	updateQuery := `UPDATE accounts SET balance = balance - $1 WHERE id = $2`
	_, err = db.DB.Exec(updateQuery, amount, accountID)
	if err != nil {
		return fmt.Errorf("withdrawal failed: %v", err)
	}

	return nil
}
