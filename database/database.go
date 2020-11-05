package database

import (
	"gormstudy/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Instance return db instance
func Instance() *gorm.DB {
	return db
}

func initDatabase() {
	dsn := "user=postgres password=docker dbname=gormstudy host=localhost port=5432 sslmode=disable"
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// Setup database
func Setup() {
	initDatabase()

	Instance().AutoMigrate(&models.User{})
	Instance().AutoMigrate(&models.Card{})
	Instance().AutoMigrate(&models.Product{})
	Instance().AutoMigrate(&models.Sale{})
	err := Instance().SetupJoinTable(&models.User{}, "Products", &models.Sale{})

	if err != nil {
		log.Fatalln(err)
	}
}
