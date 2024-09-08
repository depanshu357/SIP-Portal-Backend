package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connection established")
	}
}

func CloseDb() {
	db, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	} else {
		log.Println("Database connection closed")
	}
	db.Close()
}
