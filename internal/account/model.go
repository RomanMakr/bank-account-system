package account


// AccountRequest defines the input structure for account creation
type AccountRequest struct {
	// Name of the account holder
	// Example: Alice
	Name string `json:"name" example:"Alice"`

	// Initial balance of the account
	// Example: 500.0
	Balance float64 `json:"balance" example:"500.0"`
}


// Account represents the account data model
// @Description Account data including the account ID, account holder's name, and the balance
type Account struct {
	ID      int     `json:"id" example:"1"`       // Account ID
	Name    string  `json:"name" example:"Jon"`    // Account holder's name
	Balance float64 `json:"balance" example:"1000"` // Account balance
}
