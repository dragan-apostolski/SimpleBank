package account

import (
	"testing"
)

func TestSavingsAccountWithdraw_HappyPath(t *testing.T) {
	initBalance := 200
	withdrawalAmount := 50

	acc, err := NewSavingsAccount(1, initBalance)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = acc.Withdraw(withdrawalAmount)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if acc.Balance != initBalance-withdrawalAmount {
		t.Errorf("Withdrawal not working properly")
	}
}

func TestSavingsAccountWithdraw_WithdrawNotPossible(t *testing.T) {
	initBalance := 200
	withdrawalAmount := 110

	acc, err := NewSavingsAccount(1, initBalance)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = acc.Withdraw(withdrawalAmount)
	if err.Error() != "You can only withdraw up to 100 euros" {
		t.Errorf("Savings account must always have at least 100 euroes in balance")
	}
}

func TestSavingsAccountDeposit(t *testing.T) {
	initBalance := 200
	depositAmount := 50

	acc, err := NewSavingsAccount(1, initBalance)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = acc.Deposit(depositAmount)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if acc.Balance != initBalance+depositAmount {
		t.Errorf("Deposit on savings account not working properly")
	}
}
