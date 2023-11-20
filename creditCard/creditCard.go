package creditCard

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

func (c *CreditCard) IsValid() bool {
	return true
}
