package main

import (
	"fmt"
	"github.com/sambcox/go-credit-check/bank"
	"github.com/sambcox/go-credit-check/creditCard"
)

func main() {
	fmt.Println("Welcome to Credit Card Validator!")

	fmt.Print("Enter a credit card number: ")
	var number string
	fmt.Scanln(&number)

	fmt.Print("Enter a credit card limit: ")
	var limit int
	fmt.Scanln(&limit)

	fmt.Println("Initializing your credit card...")

	myCreditCard := creditCard.NewCreditCard(number, limit)
	b := bank.NewBank("My Bank")

	fmt.Print("Enter a transaction amount: ")
	var amount int
	fmt.Scanln(&amount)
	result := b.ValidTransaction(amount, myCreditCard)
	if result == "success" {
		fmt.Printf("The transaction for $%d on the card ending in %s is valid!", amount, myCreditCard.GetLastFour())
	} else {
		fmt.Printf("The transaction for $%d on the card ending in %s is invalid! The reason given is: %s", amount, myCreditCard.GetLastFour(), result)
	}
}
