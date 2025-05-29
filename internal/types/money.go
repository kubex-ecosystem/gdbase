package types

import (
	"fmt"
)

// Money represents a monetary value with precision to avoid floating-point rounding issues
type Money struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

// ZeroMoney represents a zero monetary value
var ZeroMoney = Money{Amount: 0, Currency: "BRL"}

// DefaultCurrency is the default currency used
const DefaultCurrency = "BRL"

// NewMoney creates a Money object from a decimal value
func NewMoney(amountDecimal float64, currency string) Money {
	if currency == "" {
		currency = DefaultCurrency
	}
	amount := int64(amountDecimal * 100)
	return Money{Amount: amount, Currency: currency}
}

// Format formats a Money object for display
func (m Money) Format() string {
	return fmt.Sprintf("%s %.2f", m.Currency, float64(m.Amount)/100)
}

// Add adds two Money values
func (m Money) Add(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, fmt.Errorf("currencies do not match: %s vs %s", m.Currency, other.Currency)
	}
	return Money{Amount: m.Amount + other.Amount, Currency: m.Currency}, nil
}

// Subtract subtracts one Money value from another
func (m Money) Subtract(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, fmt.Errorf("currencies do not match: %s vs %s", m.Currency, other.Currency)
	}
	return Money{Amount: m.Amount - other.Amount, Currency: m.Currency}, nil
}

// Multiply multiplies a Money value by a factor
func (m Money) Multiply(factor float64) Money {
	return Money{Amount: int64(float64(m.Amount) * factor), Currency: m.Currency}
}

// Divide divides a Money value by a divisor
func (m Money) Divide(divisor float64) (Money, error) {
	if divisor == 0 {
		return Money{}, fmt.Errorf("division by zero")
	}
	return Money{Amount: int64(float64(m.Amount) / divisor), Currency: m.Currency}, nil
}

// Percentage calculates a percentage of a Money value
func (m Money) Percentage(percentage float64) Money {
	return Money{Amount: int64(float64(m.Amount) * percentage / 100), Currency: m.Currency}
}

// Compare compares two Money values
// Returns -1 if m < other, 0 if m == other, 1 if m > other
func (m Money) Compare(other Money) (int, error) {
	if m.Currency != other.Currency {
		return 0, fmt.Errorf("currencies do not match: %s vs %s", m.Currency, other.Currency)
	}
	if m.Amount < other.Amount {
		return -1, nil
	} else if m.Amount > other.Amount {
		return 1, nil
	}
	return 0, nil
}

// IsNegative checks if a Money value is negative
func (m Money) IsNegative() bool {
	return m.Amount < 0
}

// IsZero checks if a Money value is zero
func (m Money) IsZero() bool {
	return m.Amount == 0
}

// Absolute returns the absolute value of a Money object
func (m Money) Absolute() Money {
	if m.Amount < 0 {
		return Money{Amount: -m.Amount, Currency: m.Currency}
	}
	return m
}
