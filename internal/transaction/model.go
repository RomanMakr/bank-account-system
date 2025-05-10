package transaction

import "time"


// TransferRequest defines the input format for a transfer operation
// @Description Transfer request with sender, receiver, and amount
type TransferRequest struct {
	// SenderID is the ID of the account sending money
	// Example: 1
	SenderID int `json:"sender_id" example:"1"`

	// ReceiverID is the ID of the account receiving money
	// Example: 2
	ReceiverID int `json:"receiver_id" example:"2"`	

	// Amount to transfer between accounts
	// Example: 150.75
	Amount float64 `json:"amount" example:"150.75"`
} 

// Transaction represents the data model for a transfer
// @Description Transfer data including sender and receiver IDs, the transferred amount, and the creation time
type Transaction struct {
	ID int 				`json:"id" example:"1"` 					// Transaction ID
	SenderID int 		`json:"sender_id" example:"1"`					// Sender's account ID
	ReceiverID int 		`json:"receiver_id" example:"2"`					// Receiver's account ID
	Amount float64 		`json:"amount" example:"100.50"`				// Amount transferred
	CreatedAt time.Time `json:"created_at" example:"2025-05-08T18:00:00Z"`	// Timestamp of the transaction
}