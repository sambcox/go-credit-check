package main

import (
	"fmt"
	"github.com/sambcox/go-credit-check/bank"
	"github.com/sambcox/go-credit-check/creditCard"
)

func main() {
	myCreditCard := creditCard.NewCreditCard("5541808923795240", 1000)
	b := bank.NewBank("My Bank")
	if b.ValidTransaction(100, myCreditCard) {
		fmt.Printf("The number %s is valid!", myCreditCard.GetNumber())
	} else {
		fmt.Printf("The number %s is invalid!", myCreditCard.GetNumber())
	}
}
