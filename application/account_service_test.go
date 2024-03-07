package application

import (
	"bank/infrastructure"
	"testing"
)

func TestAccountService_OpenSavingsAccount(t *testing.T) {
	db := infrastructure.NewEmptyDatabase()
	service := &AccountService{database: db}

	accountId := 1
	amountToDeposit := 100

	err := service.OpenSavingsAccount(accountId, amountToDeposit)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, ok := db.Accounts[accountId]
	if !ok {
		t.Errorf("Account not created in database")
	}
}

func TestAccountService_OpenCurrentAccount(t *testing.T) {
	db := infrastructure.NewEmptyDatabase()
	service := &AccountService{database: db}

	accountId := 123

	err := service.OpenCurrentAccount(accountId)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, ok := db.Accounts[accountId]
	if !ok {
		t.Errorf("Account not created in database")
	}
}

func TestAccountService_Withdraw_SavingsAccount(t *testing.T) {
	db := infrastructure.NewEmptyDatabase()
	service := &AccountService{database: db}

	accountId := 1
	balance := 160
	amountToWithdraw := 50
	_ = service.OpenSavingsAccount(accountId, balance)

	err := service.Withdraw(accountId, amountToWithdraw)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if db.Accounts[accountId].GetBalance() != balance-amountToWithdraw {
		t.Errorf("Withdrawal on savings account failed")
	}
}

func TestAccountService_Withdraw_CurrentAccount(t *testing.T) {
	db := infrastructure.NewEmptyDatabase()
	service := &AccountService{database: db}

	accountId := 1
	balance := 160
	amountToWithdraw := 50
	_ = service.OpenCurrentAccount(accountId)

	_ = service.Deposit(accountId, balance)
	err := service.Withdraw(accountId, amountToWithdraw)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if db.Accounts[accountId].GetBalance() != balance-amountToWithdraw {
		t.Errorf("Withdrawal on current account failed")
	}
}

func TestAccountService_Deposit_SavingsAccount(t *testing.T) {
	db := infrastructure.NewEmptyDatabase()
	service := &AccountService{database: db}

	accountId := 1
	balance := 160
	amountToDeposit := 50
	_ = service.OpenSavingsAccount(accountId, balance)

	err := service.Deposit(accountId, amountToDeposit)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if db.Accounts[accountId].GetBalance() != balance+amountToDeposit {
		t.Errorf("Deposit on savings account failed")
	}
}

func TestAccountService_Deposit_CurrentAccount(t *testing.T) {
	db := infrastructure.NewEmptyDatabase()
	service := &AccountService{database: db}

	accountId := 1
	amountToDeposit := 50
	_ = service.OpenCurrentAccount(accountId)

	err := service.Deposit(accountId, amountToDeposit)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if db.Accounts[accountId].GetBalance() != amountToDeposit {
		t.Errorf("Deposit on current account failed")
	}
}
