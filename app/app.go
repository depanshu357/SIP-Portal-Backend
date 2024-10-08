package app

import (
	"context"
	"net/http"
	"sip/database"
	"sip/database/migration"
	"sip/routes"
	"sip/utils"
	"time"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func Run() {

	utils.InitializeLogger()
	utils.Logger.Info("Starting the application")

	loadEnv()
	initDb()

	router := routes.InitRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func(){
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Logger.Sugar().Fatalf("Failed to run server: %v", err)
		}
	}()

	<-quit
	utils.Logger.Info("Shutting down server...")

	shutdownServer(router)

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

func shutdownServer(router *gin.Engine) {
	shutdownTimeout := 5 * time.Second
	utils.Logger.Info("Waiting for ongoing requests to finish...")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		utils.Logger.Sugar().Fatalf("Server forced to shutdown: %v", err)
	}

	utils.Logger.Info("Server exiting")
}
