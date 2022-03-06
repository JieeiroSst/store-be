package model

import "time"

type Sale struct {
	Id          string
	Amount      string
	Description string
	Type        string
	CustomerId  string
	Expire      time.Time
}

type InputSale struct {
	Amount      string
	Description string
	Type        string
	CustomerId  string
	Expire      int
}
