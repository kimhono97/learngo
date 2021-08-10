package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	//private
	owner   string
	balance int

	//public
	Memo string
}

var errNoMoney = errors.New("you can't withdraw because you are poor")

func NewAccount(owner string) *Account {
	acc := Account{owner: owner, balance: 0}
	return &acc
}

func (a *Account) Deposit(amount int) { // a is Copying Receiver
	a.balance += amount
}

func (a Account) Balance() int { // a is Referencing Receiver
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a *Account) Owner() string {
	return a.owner
}

func (a *Account) String() string { // Overriding the Method for Representive String of Struct
	return fmt.Sprint(a.owner, "'s account has $", a.balance)
}
