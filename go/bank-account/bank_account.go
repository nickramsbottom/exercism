package account

// Account at the bank
type Account struct {
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
func (a *Account) Close() (payout int64, ok bool) {
	a.open = false
	return a.balance, true
}

// Balance of the account
func (a Account) Balance() (balance int64, ok bool) {
	if !a.open {
		return 0, false
	}

	return a.balance, true
}

// Deposit change account balance
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if !a.open {
		return 0, false
	}

	if amount+a.balance < 0 {
		return a.balance, false
	}

	a.balance = amount + a.balance
	return a.balance, true
}
