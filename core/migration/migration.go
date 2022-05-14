package migration

import (
	"fmt"

	"github.com/JIeeiroSst/store/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration(config config.Config) error {
	cfg := config.Postgres
	dns := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=%v", cfg.PgDriver, cfg.PostgresqlUser,
		cfg.PostgresqlPassword, cfg.PostgresqlHost, cfg.PostgresqlPort, cfg.PostgresqlDbname,
		cfg.PostgresqlSSLMode)
	m, err := migrate.New("file://script/migrations", dns)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}

	return nil
}
