package server

import (
	"fmt"

	"github.com/JIeeiroSst/store/config"
	"github.com/JIeeiroSst/store/internal/delivery/http"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/internal/usecase"
	"github.com/JIeeiroSst/store/pkg/hash"
	"github.com/JIeeiroSst/store/pkg/jwt"
	"github.com/JIeeiroSst/store/pkg/postgres"
	"github.com/JIeeiroSst/store/pkg/redis"
	"github.com/JIeeiroSst/store/pkg/snowflake"
	"gorm.io/gorm"
)

type Server interface {
	AppServerAPI() error
}

type server struct {
	Dependency
}

type Dependency struct {
	Config       config.Config
	Repositories repository.Repositories
	Usecase      usecase.Usecase
	Http         http.Handler
	DB           *gorm.DB
}

func NewSserver(Deps Dependency) Server {
	return &server{
		Dependency: Dependency{
			Config: Deps.Config,
		},
	}
}

func (s *server) AppServerAPI() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		s.Dependency.Config.Postgres.PostgresqlHost, s.Dependency.Config.Postgres.PostgresqlUser, s.Dependency.Config.Postgres.PostgresqlPassword,
		s.Dependency.Config.Postgres.PostgresqlDbname, s.Dependency.Config.Postgres.PostgresqlPort)

	postgresConn, err := postgres.InitPostgreSQL(dsn)
	if err != nil {
		return err
	}
	snowflake := snowflake.NewSnowflake()
	redis := redis.NewDatabase(s.Config.Redis.Dns)
	hash := hash.NewHashPassword()
	jwt := jwt.NewTokenUser(s.Config.Secret.AccessSerect, s.Config.Secret.RefreshSerect, snowflake)
	repository := repository.NewRepositories(postgresConn)
	usecase := usecase.NewUsecase(usecase.Dependency{
		Repos:     repository,
		Snowflake: snowflake,
		Hash:      hash,
		Jwt:       jwt,
	})

	handler := http.NewHandler(*usecase, jwt, redis)
	handler.Init()
	return nil
}
