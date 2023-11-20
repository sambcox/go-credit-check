package bank

import (
	"github.com/sambcox/go-credit-check/creditCard"
)

type Bank struct {
	Name string
}

func NewBank(name string) Bank {
	return Bank{Name: name}
}

func (b *Bank) ValidTransaction(amount int, myCreditCard creditCard.CreditCard) bool {
	if myCreditCard.GetLimit() >= amount && myCreditCard.IsValid() {
		return true
	} else {
		return false
	}
}
