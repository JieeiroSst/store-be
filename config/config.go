package config

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

type Config struct {
	Server         ServerConfig
	Mysql          MysqlConfig
	Secret         SerectConfig
	Postgres       PostgresConfig
	Redis          RedisConfig
	StorageService StorageServiceConfig
	Logstash       LogstashConfig
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

type RedisConfig struct {
	Dns string
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

type StorageServiceConfig struct {
	Key       string
	SecretKey string
	Endpoint  string
	Protocol  string
	Bucket    string
}

type LogstashConfig struct {
	Tranport string
	Host     string
	NameApp  string
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
