package account

type CurrentAccount struct {
	UserId  int
	Balance int
	Limit   int
}

func NewCurrentAccount(userId int, balance int, limit int) (*CurrentAccount, error) {
	return &CurrentAccount{UserId: userId, Balance: balance, Limit: limit}, nil
}
func (c *CurrentAccount) Withdraw(amount int) error {
	if amount > c.maximumWithdrawAmount() {
		return &WithdrawalAmountTooLargeError{}
	}
	c.Balance = c.Balance - amount
	return nil
}
func (c *CurrentAccount) Deposit(amount int) error {
	c.Balance = c.Balance + amount
	return nil
}

func (c *CurrentAccount) GetBalance() int {
	return c.Balance
}

func (c *CurrentAccount) maximumWithdrawAmount() int {
	return c.Limit + c.Balance
}
