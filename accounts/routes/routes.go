package routes

import (
	"Tach/controllers"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router, ac *controllers.AccountController) {
	r.HandleFunc("/accounts", ac.GetAccountsHandler).Methods("GET")
	r.HandleFunc("/accounts", controllers.CreateAccountWrapper(ac)).Methods("POST")
	r.HandleFunc("/accounts/{id}", ac.GetAccountHandler).Methods("GET")
	r.HandleFunc("/accounts/{id}", ac.UpdateAccountHandler).Methods("PUT")
	r.HandleFunc("/accounts/{id}", ac.DeleteAccountHandler).Methods("DELETE")
}
