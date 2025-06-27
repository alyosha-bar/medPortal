package database

import (
	"fmt"
	"log"
	"os"

	"github.com/alyosha-bar/medPortal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// load .env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get connection string
	connStr, exists := os.LookupEnv("DB_URL")
	if !exists {
		log.Fatal("DB_URL environment variable not set")
	}

	// connect
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	// Check connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	// AutoMigrate models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Patient{},
	)
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	fmt.Println("Connected to database successfully")
}
