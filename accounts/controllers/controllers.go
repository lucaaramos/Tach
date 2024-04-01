package controllers

import (
	"Tach/models"
	"Tach/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountController struct {
	repo *repository.AccountRepository
}

func NewAccountController(repo *repository.AccountRepository) *AccountController {
	return &AccountController{repo}
}

func CreateAccountWrapper(ac *AccountController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ac.CreateAccountHandler(w, r)
	}
}

func (ac *AccountController) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "Error decoding account data", http.StatusBadRequest)
		return
	}

	err = ac.repo.CreateAccount(r.Context(), &account)
	if err != nil {
		http.Error(w, "Error creating account", http.StatusInternalServerError)
		log.Println("Error creating account:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func (ac *AccountController) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID := params["id"]

	account, err := ac.repo.GetAccountByID(r.Context(), accountID)
	if err != nil {
		http.Error(w, "Error getting account", http.StatusInternalServerError)
		log.Println("Error getting account:", err)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func (ac *AccountController) GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	accounts, err := ac.repo.GetAllAccounts(r.Context())
	if err != nil {
		http.Error(w, "Error getting accounts", http.StatusInternalServerError)
		log.Println("Error getting accounts:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func (ac *AccountController) UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["id"]
	log.Printf("Trying to update account with ID: %s\n", accountID)

	var updateData models.Account
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := ac.repo.UpdateAccount(accountID, updateData.Name, updateData.Balance, updateData.Currency)
	if err != nil {
		log.Printf("Error updating account: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Account was successfully updated"))
}

func (ac *AccountController) DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["id"]
	err := ac.repo.DeleteAccount(accountID)
	if err != nil {
		log.Printf("Error deleting account: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Account was successfully deleted"))
}
