package account_test

import (
	"testing"
	"bank-account-system/db"
	"bank-account-system/internal/account"
)

func TestWithdraw(t *testing.T) {
	err := db.InitDB()
	if err != nil {
		t.Fatalf("Could not init DB: %v", err)
	}

	accountID, err := account.CreateAccount("Roman", 1000.0)
	if err != nil {
		t.Fatalf("Error creating account: %v", err)
	}

	amount := 500.0
	err = account.Withdraw(accountID, amount)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	balance, _ := account.GetBalance(accountID)
	if balance != 500.0 {
		t.Fatalf("Expected balance 500.0, got %v", balance)
	}
}