package account

type Account interface {
	GetBalance() int
	Withdraw(amount int) error
	Deposit(amount int) error
}
