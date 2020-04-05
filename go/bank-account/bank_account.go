package account

import (
	"sync"
)

// Account is a bank account. It has a current balance which can not be negative. The closed
// boolean says whether the account is closed or not. A closed account can not hold a balance
// or take deposits.
type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

// Open an account with a balance equal to the initialDeposit argument. It returns a pointer to the
// created account. It returns nil when the initialDeposit is less than 0.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		balance: initialDeposit,
	}
}

// Close updates the balance of the account to 0 and sets the account's closed boolean to true,
// closing the account. It returns payout which is the closing balance, plus an `ok` flag indicating
// whether the account closure was successful or not. Closed accounts can not hold a balance and can
// not be closed again.
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return a.balance, false
	}

	a.closed = true
	payout = a.balance
	a.balance = 0

	return payout, true
}

// Balance obtains the current balance of the account. It returns the current balance, plus an `ok`
// flag indicating whether obtaining the balance was successful or not. If the account is closed the
// `ok` flag will be false.
func (a *Account) Balance() (balance int64, ok bool) {
	a.RLock()
	defer a.RUnlock()
	if a.closed {
		return 0, false
	}

	return a.balance, true
}

// Deposit takes a deposit amount and adds it to the account balance. It returns the new balance, plus
// an `ok` flag indicating that the deposit was successful. If the account is closed, or if the amount
// is negative and exceeds the current balance, the `ok` flag will be false.
func (a *Account) Deposit(amount int64) (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	if amount+a.balance < 0 {
		return a.balance, false
	}

	a.balance += amount
	return a.balance, true
}
