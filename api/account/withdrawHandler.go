package account

import (
	"bank-account-system/api/utils"
	"bank-account-system/internal/account"
	"encoding/json"
	"net/http"
)

// WithdrawHandler handles the withdrawal request
// @Summary Withdraws an amount from an account
// @Description Withdraws a specified amount from the account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account_id body int true "Account ID" 
// @Param amount body float64 true "Amount to withdraw" 
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse "Invalid data input"
// @Failure 500 {object} utils.ErrorResponse "Withdrawal failed"
// @Router /withdraw [post]
func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	var withdrawRequest struct {
		AccountID int 		`json:"account_id"`
		Amount    float64	`json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&withdrawRequest)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid data input")
		return
	}

	err = account.Withdraw(withdrawRequest.AccountID, withdrawRequest.Amount)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "withdrawal failed"+err.Error())
		return
	}

		successResponse := utils.SuccessResponse{
		Data: map[string]interface{}{
			"message": "Deposit successful",
		 },
		}
	
		utils.SendSucces(w, http.StatusOK, successResponse)
}