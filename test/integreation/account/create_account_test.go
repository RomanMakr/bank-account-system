package account_test

import (
	"testing"
	"bank-account-system/db"
	"bank-account-system/internal/account"
)

func TestCreateAccount(t *testing.T) {
	errroDB := db.InitDB()
	if errroDB != nil {
		t.Fatalf("Could not init DB: %v", errroDB)
	}
	
	accountID, err := account.CreateAccount("Roman", 2000.0)
	if err != nil {
		t.Fatalf("Unexpected error while creating account: %v", err)
	}

	if accountID <= 0 {
		t.Fatalf("Expected valid account ID, got %d", accountID)
	}
}