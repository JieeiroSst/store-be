package postgres

import (
	"fmt"
	"log"
	"testing"

	"github.com/JIeeiroSst/store/config"
	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	config, err := config.ReadConfig("../../config-deploy.yml")
	if err != nil {
		log.Fatal(err)
	}

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Postgres.PostgresqlHost, config.Postgres.PostgresqlUser, config.Postgres.PostgresqlPassword,
		config.Postgres.PostgresqlDbname, config.Postgres.PostgresqlPort)

	_, err = InitPostgreSQL(dns)
	assert.NoError(t, err)
}
