package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin type
type Bitcoin int

// Stringer interface
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit method
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Withdraw method
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return errors.New("insufficient funds")
	}
	w.balance -= amount
	return nil
}

// Balance method
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
