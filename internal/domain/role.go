package domain

type Role struct {
	Id          string
	Title       string
	Description string
	Permissions []Permission `gorm:"foreignKey:RoleId"`
}
