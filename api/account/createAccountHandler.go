package account


import (
	"bank-account-system/internal/account"
	"bank-account-system/api/utils"
	"encoding/json"
	"net/http"
)



// CreateAccountHandler handles the account creation request
// @Summary Create a new bank account
// @Description Create a new bank account and return account ID and details
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body account.Account true "Account details"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /accounts [post]
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var newAccount account.AccountRequest

	err := json.NewDecoder(r.Body).Decode(&newAccount)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid data input")
		return
	}

	accountID, err := account.CreateAccount(newAccount.Name, newAccount.Balance)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create account")
		return
	}

	successResponse := utils.SuccessResponse{
		Data: map[string]interface{}{
			"account_id": accountID,
			"name":       newAccount.Name,
			"balance":    newAccount.Balance,
		},
	}

	utils.SendSucces(w, http.StatusCreated, successResponse)
}