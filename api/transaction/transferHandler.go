package transaction

import (
	"encoding/json"
	"net/http"

	"bank-account-system/api/utils"
	"bank-account-system/internal/transaction"
)

// TransferHandler handles the transfer request between accounts
// @Summary Transfers funds between two accounts
// @Description Transfers a specified amount from one account to another
// @Tags transactions
// @Accept json
// @Produce json
// @Param sender_id body int true "Sender Account ID"
// @Param receiver_id body int true "Receiver Account ID"
// @Param amount body float64 true "Amount to transfer"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse "Invalid data input"
// @Failure 500 {object} utils.ErrorResponse "Transfer failed"
// @Router /transfer [post]
func TransferHandler(w http.ResponseWriter, r *http.Request) {
	var transferRequest transaction.TransferRequest

	err := json.NewDecoder(r.Body).Decode(&transferRequest)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid data input")
		return
	}

	err = transaction.Transfer(transferRequest.SenderID, transferRequest.ReceiverID, transferRequest.Amount)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Transfer failed"+ err.Error())
		return
	}

	successResponse := utils.SuccessResponse{
		Data: map[string]interface{}{
			"message": "Deposit successful",
		 },
		}
	
		utils.SendSucces(w, http.StatusOK, successResponse)
}