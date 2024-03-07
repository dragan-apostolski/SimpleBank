package application

import (
	"bank/domain/account"
	"bank/infrastructure"
)

type AccountNotFoundException struct{}

func (a *AccountNotFoundException) Error() string {
	return "Account not found"
}

type AccountService struct {
	database *infrastructure.Database
}

func NewAccountService(database *infrastructure.Database) *AccountService {
	return &AccountService{database: database}
}

func (as *AccountService) OpenSavingsAccount(accountId int, amountToDeposit int) error {
	acc, err := account.NewSavingsAccount(accountId, amountToDeposit)
	if err != nil {
		return err
	}
	as.database.Accounts[accountId] = acc
	return nil
}

func (as *AccountService) OpenCurrentAccount(accountId int) error {
	acc, err := account.NewCurrentAccount(accountId, 0, 0)
	if err != nil {
		return err
	}
	as.database.Accounts[accountId] = acc
	return nil
}

func (as *AccountService) Withdraw(accountId int, amountToWithdraw int) error {
	acc, ok := as.database.Accounts[accountId]
	if !ok {
		return &AccountNotFoundException{}
	}

	err := acc.Withdraw(amountToWithdraw)
	if err != nil {
		return err
	}

	return nil
}

func (as *AccountService) Deposit(accountId int, amountToDeposit int) error {
	acc, ok := as.database.Accounts[accountId]
	if !ok {
		return &AccountNotFoundException{}
	}

	err := acc.Deposit(amountToDeposit)
	if err != nil {
		return err
	}

	return nil
}

func (as *AccountService) GetBalance(accountId int) (*int, error) {
	acc, ok := as.database.Accounts[accountId]
	if !ok {
		return nil, &AccountNotFoundException{}
	}

	balance, err := acc.GetBalance()
	if err != nil {
		return nil, err
	}

	return &balance, nil
}
