package routes

import (
	"transactions/controllers"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router, tc *controllers.TransactionController) {
	r.HandleFunc("/transactions", tc.GetAllTransactionsHandler).Methods("GET")
	r.HandleFunc("/transactions", tc.CreateTransactionHandler).Methods("POST")
	r.HandleFunc("/transactions/{id}", tc.GetTransactionHandler).Methods("GET")
	r.HandleFunc("/transactions/{id}", tc.UpdateTransactionHandler).Methods("PUT")
	r.HandleFunc("/transactions/{id}", tc.DeleteTransactionHandler).Methods("DELETE")
}
