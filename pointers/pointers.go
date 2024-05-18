package pointers

import "fmt"

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

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
