package migration

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration(dns string) error {
	//"postgres://postgres:postgres@localhost:5432/example?sslmode=disable"
	m, err := migrate.New("file://script/migrations",dns)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}

	return nil
}
