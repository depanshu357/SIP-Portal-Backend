package app

import (
	"sip/database"
	"sip/database/migration"
	"sip/routes"
	"sip/utils"

	"github.com/lpernett/godotenv"
)

func Run() {
	utils.InitializeLogger()
	utils.Logger.Info("Starting the application")
	loadEnv()
	initDb()
	router := routes.InitRoutes()

	router.Run(":8080")

	defer closeDb()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		utils.Logger.Fatal("Error loading .env file")
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
