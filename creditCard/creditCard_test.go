package creditCard

import (
	"testing"
)

func TestCreditCard_GetLimit(t *testing.T) {
	card := NewCreditCard("1234567890123456", 1000)
	limit := card.GetLimit()

	if limit != 1000 {
		t.Errorf("Expected limit to be 1000, but got %d", limit)
	}
}

func TestCreditCard_GetNumber(t *testing.T) {
	card := NewCreditCard("1234567890123456", 1000)
	number := card.GetNumber()

	if number != "1234567890123456" {
		t.Errorf("Expected number to be 1234567890123456, but got %s", number)
	}
}

func TestCreditCard_GetLastFour(t *testing.T) {
	card := NewCreditCard("1234567890123456", 1000)
	lastFour := card.GetLastFour()

	if lastFour != "3456" {
		t.Errorf("Expected last four digits to be 3456, but got %s", lastFour)
	}
}

func TestCreditCard_IsValid(t *testing.T) {
	card := NewCreditCard("5541808923795240", 1000)
	isValid := card.IsValid()

	if !isValid {
		t.Errorf("Expected card to be valid, but got %t", isValid)
	}

	card = NewCreditCard("4024007136512380", 1000)
	isValid = card.IsValid()

	if !isValid {
		t.Errorf("Expected card is valid to be true, but got %t", isValid)
	}

	card = NewCreditCard("5541801923795240", 1000)
	isValid = card.IsValid()

	if isValid {
		t.Errorf("Expected card is valid to be false, but got %t", isValid)
	}

	card = NewCreditCard("6011797668868728", 1000)
	isValid = card.IsValid()

	if isValid {
		t.Errorf("Expected card is valid to be false, but got %t", isValid)
	}
}
