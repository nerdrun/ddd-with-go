package domain

import "fmt"

type Price struct {
	amount   int
	currency string
}

func NewPrice(amount int, currency string) (*Price, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("price must be greater than 0")
	}
	if currency == "" {
		return nil, fmt.Errorf("currency cannot be empty")
	}
	return &Price{
		amount:   amount,
		currency: currency,
	}, nil
}
