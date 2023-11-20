package bank

import (
	"github.com/sambcox/go-credit-check/creditCard"
	"testing"
)

func TestBank_ValidTransaction(t *testing.T) {
	card := creditCard.NewCreditCard("5541808923795240", 1000)
	bank := NewBank("My Bank")
	isValid := bank.ValidTransaction(100, card)

	if !isValid {
		t.Errorf("Expected transaction to be valid, but got %t", isValid)
	}

	card = creditCard.NewCreditCard("6011797668868728", 1000)
	isValid = bank.ValidTransaction(100, card)

	if isValid {
		t.Errorf("Expected transaction to be invalid, but got %t", isValid)
	}

	card = creditCard.NewCreditCard("5541808923795240", 1000)
	isValid = bank.ValidTransaction(10000, card)

	if isValid {
		t.Errorf("Expected transaction to be invalid, but got %t", isValid)
	}
}
