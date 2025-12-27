package expense

import "time"

type Expense struct {
	ID       int
	Title     string
	Amount   float64
	Category string
	Date     time.Time
}

type ExpenseList []Expense