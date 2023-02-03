package database

import (
	"ecom/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}
	
	
	log.Println("Successfully connected to DB")

	db.Logger = logger.Default.LogMode(logger.Info)

	//Migrations

	log.Println("Running Migrations")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
		
	Database = DbInstance{Db: db}

}