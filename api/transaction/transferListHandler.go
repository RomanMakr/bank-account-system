package transaction

import (
	"encoding/json"
	"net/http"

	"bank-account-system/api/utils"
	"bank-account-system/db"
)

const listTransferQuery = `
	SELECT id, sender_id, receiver_id, amount, created_at
	FROM transactions
	ORDER BY created_at DESC
`

// Transaction represents a transaction record in the system
// @Description A transaction includes details of the sender, receiver, amount, and the timestamp of the transfer
type Transaction struct {
	// ID is the unique identifier of the transaction
	// Example: 1
	ID 			int 	`json:"id" example:"1"`
	
	// SenderID is the ID of the sender account
	// Example: 1
	SenderID 	int 	`json:"sender_id" example:"1"`
	
	// ReceiverID is the ID of the receiver account
	// Example: 2
	ReceiverID 	int 	`json:"receiver_id" example:"2"`
	
	// Amount is the amount of money transferred
	// Example: 150.75	
	Amount 		float64	`json:"amount" example:"150.75"`
	
	// CreatedAt is the timestamp when the transaction was created
	// Example: "2025-05-08T12:30:00Z"
	CreatedAt 	string 	`json:"created_at" example:"2025-05-08T12:30:00Z"`
}


// TransferListHandler handles the GET request to list all transactions
// @Summary List all transactions
// @Description Returns a list of all money transfers sorted by most recent first
// @Tags Transactions
// @Produce json
// @Success 200 {array} transaction.Transaction
// @Failure 500 {object} utils.ErrorResponse
// @Router /transactions [get]

func TransferListHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(listTransferQuery)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetch transaction")
		return
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.ID, &transaction.SenderID, &transaction.ReceiverID, &transaction.Amount, &transaction.CreatedAt)
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, "Failed to scan transaction")
			return
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to iterate over transactions")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}