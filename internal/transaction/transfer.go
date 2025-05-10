package transaction

import (
	"fmt"
	"bank-account-system/db"
)

func Transfer(senderID, receiverID int, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("the amount must be positive")
	}

	
	transaction, err := db.DB.Begin()
	if err != nil {
		return fmt.Errorf("could not start transaction: %v", err)
	}
	
	var balance float64

	err = transaction.QueryRow(`SELECT balance FROM accounts WHERE id = $1 FOR UPDATE`, senderID).Scan(&balance)
	if err != nil {
		transaction.Rollback()
		return fmt.Errorf("could not fetch sender's account balance: %v", err)
	}

	if balance < amount {
		transaction.Rollback()
		return fmt.Errorf("insufficient funds")
	}


	updateSenderQuery := `UPDATE accounts SET balance = balance - $1 WHERE id = $2`
	_, err = transaction.Exec(updateSenderQuery, amount, senderID)
	if err != nil {
		transaction.Rollback()
		return fmt.Errorf("could not update sender's account: %v", err)
	}

	updateReceiverQuery := `UPDATE accounts SET balance = balance + $1 WHERE id = $2`
	_, err = transaction.Exec(updateReceiverQuery, amount, receiverID)
	if err != nil {
		transaction.Rollback()
		return fmt.Errorf("could not update receiver's account: %v", err)
	}

	insertTransactionQuery := `INSERT INTO transactions (sender_id, receiver_id, amount) VALUES ($1, $2, $3)`
	_, err = transaction.Exec(insertTransactionQuery, senderID, receiverID, amount)
	if err != nil {
		transaction.Rollback()
		return fmt.Errorf("could not insert transaction record: %v", err)
	}

	err = transaction.Commit()
	if err != nil {
		return fmt.Errorf("could not commit transaction: %v", err)
	}

	return nil
}