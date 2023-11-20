# Credit Card Validator

Welcome to the Credit Card Validator project! This simple Go application allows you to validate credit card transactions against a specified limit. This project was originally created for the Turing School of Software and Design's Back End Engineering program in Ruby, but I have adapted the original project, while following all guidlines, to work in Go.

## Usage

1. Clone the repository:

    ```bash
    git clone https://github.com/sambcox/go-credit-check.git
    ```

1. Navigate to the project directory:

    ```bash
    cd go-credit-check
    ```

1. Run the application:

    ```bash
    go run main.go
    ```

1. Confirm that all tests are passing:

    ```bash
    go test ./...
    ```

1. Follow the prompts in the command line to input credit card details and transaction amount.


## Project Structure

* **main.go**: Contains the main application logic for interacting with users and initiating credit card validations.
* **creditCard/creditCard.go**: Defines the **CreditCard** struct and related methods for creating, retrieving information, and validating credit cards.
* **bank/bank.go**: Implements the **Bank** struct and a method for validating transactions based on credit card limits and validity.