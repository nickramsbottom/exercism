package account

import (
	"sync"
)

// Account at the bank
type Account struct {
	sync.Mutex
	balance int64
	open    bool
}

// Open an account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		balance: initialDeposit,
		open:    true,
	}
}

// Close the account
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return a.balance, false
	}

	a.open = false
	payout := a.balance
	a.balance = 0

	return payout, true
}

// Balance of the account
func (a *Account) Balance() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open {
		return 0, false
	}

	return a.balance, true
}

// Deposit change account balance
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.open {
		return 0, false
	}

	if amount+a.balance < 0 {
		return a.balance, false
	}

	a.balance = amount + a.balance
	return a.balance, true
}
