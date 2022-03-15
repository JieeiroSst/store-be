package model

import "time"

type User struct {
	Id          string
	RoleId      string
	Username    string
	Password    string
	Email       string
	Name        string
	Description string
	Address     string
	CreatedAt   time.Time
	Lock        int
	Carts       []Cart  `gorm:"foreignKey:UserId"`
	Roles       []Role  `gorm:"foreignKey:Id;references:RoleId"`
	Sale        Sale    `gorm:"foreignKey:CustomerId"`
	Payment     Payment `gorm:"foreignKey:CustomerId"`
}

type UserDetail struct {
	Id        string
	Username  string
	CreatedAt time.Time
	Lock      int
}
