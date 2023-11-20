package creditCard

import (
	"fmt"
	"strconv"
)

type CreditCard struct {
	Number string
	Limit  int
}

func NewCreditCard(number string, limit int) CreditCard {
	return CreditCard{Number: number, Limit: limit}
}

func (c *CreditCard) GetLimit() int {
	return c.Limit
}

func (c *CreditCard) GetNumber() string {
	return c.Number
}

func (c *CreditCard) GetLastFour() string {
	return c.Number[len(c.Number)-4:]
}

func (c *CreditCard) IsValid() bool {
	result_num := 0
	for i, char := range c.Number {

		indv_num, err := strconv.Atoi(string(char))

		if err != nil {
			fmt.Println(err)
			return false
		}

		if i%2 == 0 {
			indv_num = indv_num * 2
			if indv_num > 9 {
				indv_num = indv_num - 9
			}
		}
		result_num = result_num + indv_num
	}

	return result_num%10 == 0
}
