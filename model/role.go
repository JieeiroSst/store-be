package model

type Role struct {
	Id          string
	Title       string
	Description string
}

type InputRole struct {
	Title       string
	Description string
}
