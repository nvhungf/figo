package entity

import "time"

// Expense type
type Expense struct {
	ID         string
	Title      string
	Total      float64
	Attachment string
	CreatedAt  time.Time
}
