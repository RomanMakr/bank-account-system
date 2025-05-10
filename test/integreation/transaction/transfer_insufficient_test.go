package transaction_test


import (
	"bank-account-system/db"
	"bank-account-system/internal/transaction"
	"testing"
)

func TestTransferInsufficient(t *testing.T) {
	errorDB := db.InitDB()
	if errorDB != nil {
		t.Fatalf("Failed to InitDB: %v", errorDB)
	}

	var senderID, receiverID int

	err := db.DB.QueryRow(
		`INSERT INTO accounts (name, balance) VALUES ($1, $2) RETURNING id`,
		"Account1", 2000.0,
	).Scan(&senderID)
	if err != nil {
		t.Fatalf("Could not create sender account: %v", err)
	}

	err = db.DB.QueryRow(
		`INSERT INTO accounts (name, balance) VALUES ($1, $2) RETURNING id`,
		"Account2", 3000.0,
	).Scan(&receiverID)
	if err != nil {
		t.Fatalf("Could not create receiver account: %v", err)
	}

	defer func ()  {
		_, err := db.DB.Exec(`DELETE FROM transactions WHERE sender_id = $1 OR receiver_id = $2`, senderID, receiverID)
		if err != nil {
			t.Fatalf("Could not clean up transactions: %v", err)
		}
		
		_, err = db.DB.Exec(`DELETE FROM accounts WHERE id IN ($1, $2)`, senderID, receiverID)
		if err != nil {
			t.Fatalf("Could not clean up test account: %v", err)
		}
	}()

	err = transaction.Transfer(senderID, receiverID, 3000.0)
	if err == nil {
		t.Fatalf("Expected error due to insufficient funds, but got none")
	}

	var senderBalance, receiverBalance float64

	err = db.DB.QueryRow(`SELECT balance FROM accounts WHERE id = $1`, senderID).Scan(&senderBalance)
	if err != nil {
		t.Fatalf("Could not fetch sender's balance: %v", err)
	}

	err = db.DB.QueryRow(`SELECT balance FROM accounts WHERE id = $1`, receiverID).Scan(&receiverBalance)
	if err != nil {
		t.Fatalf("Could not fetch receiver's balance: %v", err)
	}

	if senderBalance != 2000.0 {
		t.Fatalf("Expected sender's balance 2000.0, got %v", senderBalance)
	}

	if receiverBalance != 3000.0 {
		t.Fatalf("Expected receiver's balance 3000.0, got %v", receiverBalance)
	}	
}