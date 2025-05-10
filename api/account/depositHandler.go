package account

import (
	"bank-account-system/api/utils"
	"bank-account-system/internal/account"
	"encoding/json"
	"net/http"
)

// DepositHandler handles the deposit request to an account
// @Summary Deposit money into an account
// @Description Deposit a specified amount of money into the given account ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param depositRequest body struct {
//    AccountID int     `json:"account_id"`
//    Amount    float64 `json:"amount"`
// } true "Deposit request"
// @Success 200 {object} utils.SuccessResponse{"data": {"message": "Deposit successful"}}
// @Failure 400 {object} utils.ErrorResponse{"error": "Invalid data input"}
// @Failure 500 {object} utils.ErrorResponse{"error": "Deposit failed: <error_message>"}
// @Router /accounts/deposit [post]

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	var depositRequest struct {
		AccountID int 		`json:"account_id"`
		Amount 	  float64	`json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&depositRequest)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid data input")
		return
	}

	err = account.Deposit(depositRequest.AccountID, depositRequest.Amount)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Deposit failed:"+err.Error())
		return
	}

	successResponse := utils.SuccessResponse{
	Data: map[string]interface{}{
		"message": "Deposit successful",
	 },
	}

	utils.SendSucces(w, http.StatusOK, successResponse)
}
