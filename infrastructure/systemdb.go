package infrastructure

import "bank/domain/account"

type Database struct {
	Accounts map[int]account.Account
}

func NewDatabase() *Database {
	systemDb := &Database{
		Accounts: make(map[int]account.Account),
	}
	// Pre-populate accounts
	savings1, _ := account.NewSavingsAccount(1, 200)
	savings2, _ := account.NewSavingsAccount(2, 500)
	current1, _ := account.NewCurrentAccount(3, 100, 1000)
	current2, _ := account.NewCurrentAccount(4, -500, 2000)

	systemDb.Accounts[1] = savings1
	systemDb.Accounts[2] = savings2
	systemDb.Accounts[3] = current1
	systemDb.Accounts[4] = current2
	return systemDb
}

func NewEmptyDatabase() *Database {
	systemDb := &Database{
		Accounts: make(map[int]account.Account),
	}
	return systemDb
}
