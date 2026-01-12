package main

import "time"

// Expense represents a single expense record
type Expense struct {
	ID            int
	Amount        float64
	Description   string
	Category      string
	Date          time.Time
	PaymentMethod string // "Cash" or "UPI"
	Notes         string
}
