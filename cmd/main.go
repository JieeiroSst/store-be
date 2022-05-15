package main

import (
	"log"

	"github.com/JIeeiroSst/store/config"
	"github.com/JIeeiroSst/store/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	nameFile := "config.yml"

	config, err := config.ReadConfig(nameFile)
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(*config)
	if err := app.RunApp(*router); err != nil {
		log.Println(err)
	}

	router.Run(config.Server.PortServer)
}
