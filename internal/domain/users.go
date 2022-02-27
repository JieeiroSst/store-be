package domain

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

type Token struct {
	Id       string
	RoleId   string
	Username string
	Password string
}

type InputUser struct {
	RoleId      string
	Username    string
	Password    string
	Email       string
	Name        string
	Description string
	Address     string
	CreatedAt   time.Time
	Lock        int
}
