package domain

type CasbinRule struct {
	ID    int `gorm:"primaryKey;autoIncrement"`
	Ptype string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}
