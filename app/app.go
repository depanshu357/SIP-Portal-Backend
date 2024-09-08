package app

import (
	"log"
	"sip/database"
	"sip/database/migration"
	"sip/routes"

	"github.com/lpernett/godotenv"
)

func Run() {
	loadEnv()
	initDb()
	router := routes.InitRoutes()

	router.Run(":8080")

	defer closeDb()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func initDb() {
	database.ConnectToDb()
	migration.Up()
	// defer database.CloseDb()
}

func closeDb() {
	database.CloseDb()
}
