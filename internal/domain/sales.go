package domain

import "time"

type Sale struct {
	Id          string
	Amount      string
	Description string
	Type        string
	CustomerId  string
	Expire      time.Time
}

type Expire struct {
	Expire time.Time
}
