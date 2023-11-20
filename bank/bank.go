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

func (b *Bank) ValidTransaction(amount int, myCreditCard creditCard.CreditCard) string {
	if myCreditCard.GetLimit() >= amount && myCreditCard.IsValid() {
		return "success"
	} else {
		if myCreditCard.GetLimit() < amount {
			return "Credit limit exceeded"
		} else {
			return "Card is invalid"
		}
	}
}
