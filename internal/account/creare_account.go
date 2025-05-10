package account

import (
	"bank-account-system/db"
	"fmt"
)

func CreateAccount(name string, balance float64) (int, error) {
	query := `INSERT INTO accounts (name, balance) VALUES ($1, $2) RETURNING id`
	
   	var id int 
	
	err := db.DB.QueryRow(query, name, balance).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("account creation failed: %v", err)
	}

	return id, nil
}