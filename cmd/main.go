package main

import (
	"log"

	"github.com/JIeeiroSst/store/config"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config, err := config.ReadConfig("config-local.yml")
	if err != nil {
		log.Fatal(err)
	}

	router.Run(config.Server.PortServer)
}
