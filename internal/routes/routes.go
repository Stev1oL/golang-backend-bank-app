package routes

import (
	"net/http"

	"github.com/steviol/golang-backend-bank-app/internal/delivery/handler"
)

func SetupRoutes(accountHandler *handler.AccountHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/accounts", accountHandler.CreateAccount)
	mux.HandleFunc("/accounts/balance", accountHandler.AddBalance)
	mux.HandleFunc("/accounts/transactions", accountHandler.GetTransactions)

	mux.HandleFunc("/transfer", accountHandler.Transfer)

	mux.HandleFunc("/qr/generate", accountHandler.GenerateQRCode)
	mux.HandleFunc("/qr/process", accountHandler.ProcessQRPayment)

	return mux
}
