package app

import (
	"strings"

	"github.com/JIeeiroSst/store/config"
	"github.com/JIeeiroSst/store/internal/server"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type App interface {
	RunApp(engine gin.Engine) error
}

type app struct {
	Config config.Config
}

func NewApp(config config.Config) App {
	return &app{
		Config: config,
	}
}

func (a *app) RunApp(engine gin.Engine) error {
	p := ginprometheus.NewPrometheus("gin")

	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			if p.Key == "name" {
				url = strings.Replace(url, p.Value, ":name", 1)
				break
			}
		}
		return url
	}

	p.Use(&engine)
	server := server.NewServer(server.Dependency{
		Config: a.Config,
	})

	if err := server.AppServerAPI(); err != nil {
		return err
	}

	return nil

}
