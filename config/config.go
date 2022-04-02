package config

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

type Config struct {
	Server   ServerConfig
	Mysql    MysqlConfig
	Secret   SerectConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	PortServer    string
	PprofPortGRPC string
	PprofPortHTTP string
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

type MysqlConfig struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDbname   string
	MysqlSSLMode  bool
	MysqlDriver   string
}

type SerectConfig struct {
	JwtSecretKey  string
	OtpSecretKey  string
	AccessSerect  string
	RefreshSerect string
}

func ReadConfig(filename string) (*Config, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return config, nil
}
