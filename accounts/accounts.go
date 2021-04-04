package accounts

import (
	"errors"
	"fmt"
)

var errNotEnoughMoney = errors.New("Not Enough Money")

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount return pointer to avoid copy struct
func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 0}
}

// Owner getter
func (a Account) Owner() string {
	return a.owner
}

// Balance getter
func (a Account) Balance() int {
	return a.balance
}

// Withdraw balance
// pointer receiver to modify struct value
func (a *Account) Withdraw(money int) error {
	if a.balance < money {
		return errNotEnoughMoney
	}
	a.balance -= money
	return nil
}

// Deposit balance
func (a *Account) Deposit(money int) {
	a.balance += money
}

// String override like override toString()
func (a Account) String() string {
	return fmt.Sprint("owner : ", a.owner, ", balance : ", a.balance)
}
