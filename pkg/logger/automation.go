package logger

import (
	"log"
	"os"
)

func NewAutomationLogger() {
	db, closeDB, err := OpenDB("localhost", 5432, "logging_golang")
	if err != nil {
		log.Println(err)
	}
	defer closeDB()
	app := SetupCLI(db)
	if err := app.Run(os.Args); err != nil {
		log.Println(err)
	}
}
