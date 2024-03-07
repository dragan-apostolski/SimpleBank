package presentation

import (
	"bank/application"
	"log"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	service *application.AccountService
}

func NewAccountHandler(service *application.AccountService) *AccountHandler {
	return &AccountHandler{service}
}

func (ah *AccountHandler) OpenSavingsAccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	accountID, err := strconv.Atoi(r.FormValue("account_id"))
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	amountToDeposit, err := strconv.Atoi(r.FormValue("amount_to_deposit"))
	if err != nil {
		http.Error(w, "Invalid amount to deposit", http.StatusBadRequest)
		return
	}

	err = ah.service.OpenSavingsAccount(accountID, amountToDeposit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ah *AccountHandler) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	accountID, err := strconv.Atoi(r.FormValue("account_id"))
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	amountToWithdraw, err := strconv.Atoi(r.FormValue("amount_to_withdraw"))
	if err != nil {
		http.Error(w, "Invalid amount to withdraw", http.StatusBadRequest)
		return
	}

	err = ah.service.Withdraw(accountID, amountToWithdraw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ah *AccountHandler) DepositHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	accountID, err := strconv.Atoi(r.FormValue("account_id"))
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	amountToDeposit, err := strconv.Atoi(r.FormValue("amount_to_deposit"))
	if err != nil {
		http.Error(w, "Invalid amount to deposit", http.StatusBadRequest)
		return
	}

	err = ah.service.Deposit(accountID, amountToDeposit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func StartHTTPServer(service *application.AccountService) {
	handler := NewAccountHandler(service)

	http.HandleFunc("/open-savings-account", handler.OpenSavingsAccountHandler)
	http.HandleFunc("/withdraw", handler.WithdrawHandler)
	http.HandleFunc("/deposit", handler.DepositHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
