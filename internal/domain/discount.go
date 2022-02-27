package domain

import "time"

type Discount struct {
	Id          string
	Name        string
	Description string
	Type        string
	ProductId   string
	StartTime   time.Time `json:"startTime" form:"startTime" time_format:"2006-01-02T15:04:05Z07:00"`
	EndTime     time.Time `json:"endTime" form:"endTime" time_format:"2006-01-02T15:04:05Z07:00"`
}
