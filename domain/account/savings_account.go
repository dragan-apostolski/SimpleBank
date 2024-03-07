package account

type SavingsAccount struct {
	UserId  int
	Balance int
}

func NewSavingsAccount(userId int, balance int) (*SavingsAccount, error) {
	if balance < 100 {
		return nil, &TooLowDepositAmountError{}
	}
	return &SavingsAccount{UserId: userId, Balance: balance}, nil
}
func (s *SavingsAccount) Withdraw(amount int) error {
	if s.Balance-amount < 100 {
		return &WithdrawNotPossible{MaximumWithdrawAmountPossible: s.Balance - 100}
	}
	s.Balance = s.Balance - amount
	return nil
}
func (s *SavingsAccount) Deposit(amount int) error {
	s.Balance = s.Balance + amount
	return nil
}

func (s *SavingsAccount) GetBalance() (int, error) {
	return s.Balance, nil
}
