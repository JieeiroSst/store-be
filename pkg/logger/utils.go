package logger

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(host string, port int, dbname string) (*sql.DB, func() error, error) {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test?sql_mode=TRADITIONAL")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if err != nil {
		return nil, nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, nil, err
	}

	return db, db.Close, err
}
