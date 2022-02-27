package domain

import "time"

type Payment struct {
	Id          string
	CustomerId  string
	Amount      string
	Date        time.Time
	Description string
}
