package account

import (
	"testing"
)

func TestCurrentAccountWithdraw_HappyPath(t *testing.T) {
	initBalance := 100
	limit := 1000
	withdrawalAmount := 150

	acc, err := NewCurrentAccount(1, initBalance, limit)
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

func TestCurrentAccountWithdrawal_AmountTooLarge(t *testing.T) {
	initBalance := 100
	limit := 1000
	withdrawalAmount := 1500

	acc, err := NewCurrentAccount(1, initBalance, limit)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = acc.Withdraw(withdrawalAmount)
	if err.Error() != "Withdrawal amount is too large" {
		t.Errorf("Check for maximum withdrawal amount not working")
	}
}

func TestCurrentAccountDeposit(t *testing.T) {
	initBalance := 100
	limit := 1000
	depositAmount := 150

	acc, err := NewCurrentAccount(1, initBalance, limit)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = acc.Deposit(depositAmount)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if acc.Balance != initBalance+depositAmount {
		t.Errorf("Deposit on current account not working properly")
	}
}
