package server

import (
	"github.com/JIeeiroSst/store/internal/delivery/http"
	"github.com/JIeeiroSst/store/internal/repository"
	"github.com/JIeeiroSst/store/internal/usecase"
)

type Server interface {
	AppServerAPI() error
}

type server struct {
	Dependency
}

type Dependency struct {
	Repositories repository.Repositories
	Usecase      usecase.Usecase
	Http         http.Handler
}

func NewSserver(Deps Dependency) Server {
	return &server{
		Dependency: Dependency{
			Repositories: Deps.Repositories,
			Usecase:      Deps.Usecase,
			Http:         Deps.Http,
		},
	}
}

func (s *server) AppServerAPI() error {

	return nil
}
