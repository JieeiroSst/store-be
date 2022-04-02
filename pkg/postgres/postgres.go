package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *PostgresConnect
)

type PostgresConnect struct {
	db *gorm.DB
}

func InitPostgreSQL(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	instance = &PostgresConnect{db}

	return db, nil
}
