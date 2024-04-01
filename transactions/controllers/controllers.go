package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"transactions/models"
	"transactions/repository"

	"github.com/gorilla/mux"
)

type TransactionController struct {
	repo *repository.TransactionRepository
}

func NewTransactionController(repo *repository.TransactionRepository) *TransactionController {
	return &TransactionController{repo}
}

func (tc *TransactionController) GetAllTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	transactions, err := tc.repo.GetAllTransactions(r.Context())
	if err != nil {
		http.Error(w, "Error getting all transactions", http.StatusInternalServerError)
		log.Println("Error getting all transactions:", err)
		return
	}

	json.NewEncoder(w).Encode(transactions)
}

func (tc *TransactionController) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := tc.repo.CreateTransaction(r.Context(), &transaction); err != nil {
		log.Println("Error creating transaction:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "The transaction was created successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (tc *TransactionController) GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionID := params["id"]

	transaction, err := tc.repo.GetTransactionByID(r.Context(), transactionID)
	if err != nil {
		http.Error(w, "Error getting transaction", http.StatusInternalServerError)
		log.Println("Error getting transaction:", err)
		return
	}
	json.NewEncoder(w).Encode(transaction)
}

func (tc *TransactionController) UpdateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID := vars["id"]
	log.Printf("Trying to update transaction with ID: %s\n", transactionID)

	var updatedTransaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&updatedTransaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := tc.repo.UpdateTransaction(transactionID, &updatedTransaction)
	if err != nil {
		log.Printf("Error updating transaction: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "The transaction has been updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (tc *TransactionController) DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionID := params["id"]

	err := tc.repo.DeleteTransaction(r.Context(), transactionID)
	if err != nil {
		http.Error(w, "Error deleting transaction", http.StatusInternalServerError)
		log.Println("Error deleting transaction:", err)
		return
	}
	message := "Transaction deleted successfully"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
