package app

import (
	"github.com/JIeeiroSst/store/config"
	"github.com/JIeeiroSst/store/internal/server"
)

type App interface {
	RunApp() error
}

type app struct {
	Config config.Config
}

func NewApp(config config.Config) App {
	return &app{
		Config: config,
	}
}

func (a *app) RunApp() error {
	server := server.NewServer(server.Dependency{
		Config: a.Config,
	})

	if err := server.AppServerAPI(); err != nil {
		return err
	}

	return nil

}
