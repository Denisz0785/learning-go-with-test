package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposite(amount Bitcoin) {
	w.balance += amount

}

func (b *Wallet) Balance() Bitcoin {
	return b.balance
}

var ErrInsufficientFound = errors.New("not enough money")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if (w.balance - amount) < 0 {
		return ErrInsufficientFound
	}
	w.balance -= amount
	return nil
}
