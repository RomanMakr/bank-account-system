package api

import (
	"bank-account-system/api/account"
	"bank-account-system/api/transaction"
	_ "bank-account-system/docs"		
	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/accounts", account.CreateAccountHandler).Methods("POST")
	router.HandleFunc("/accounts/deposit", account.DepositHandler).Methods("POST")
	router.HandleFunc("/accounts/withdraw", account.WithdrawHandler).Methods("POST")

	router.HandleFunc("/transfer", transaction.TransferHandler).Methods("POST")
	router.HandleFunc("/transfer/list", transaction.TransferListHandler).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)	
	return router
}