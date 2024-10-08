package app

import (
	"sip/database"
	"sip/database/migration"
	"sip/routes"
	"sip/utils"
	"sync"

	"github.com/lpernett/godotenv"
)

func Run(wg *sync.WaitGroup) {
	utils.InitializeLogger()
	utils.Logger.Info("Starting the application")
	loadEnv()
	initDb()
	router := routes.InitRoutes()
	router.Run(":8080")
	defer wg.Done()
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
