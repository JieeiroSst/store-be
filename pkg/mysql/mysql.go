package mysql

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *MysqlConnect
	once     sync.Once
)

type MysqlConnect struct {
	db *gorm.DB
}

func InitMysql(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	instance = &MysqlConnect{db: db}

	return db, nil
}
