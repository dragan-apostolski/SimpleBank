package account

import "fmt"

type WithdrawalAmountTooLargeError struct{}

func (w *WithdrawalAmountTooLargeError) Error() string {
	return "Withdrawal amount is too large"
}

type TooLowDepositAmountError struct{}

func (t *TooLowDepositAmountError) Error() string {
	return "The deposit amount cannot lower than 100"
}

type WithdrawNotPossible struct {
	MaximumWithdrawAmountPossible int
}

func (w *WithdrawNotPossible) Error() string {
	return fmt.Sprintf("You can only withdraw up to %d euros", w.MaximumWithdrawAmountPossible)
}
