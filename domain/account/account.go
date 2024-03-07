package account

type Account interface {
	GetBalance() (int, error)
	Withdraw(amount int) error
	Deposit(amount int) error
}
