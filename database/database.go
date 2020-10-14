package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDatabase : init database
func InitDatabase() {
	dsn := "user=postgres password=docker dbname=gormstudy host=localhost port=5432 sslmode=disable"
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// Instance : return db instance
func Instance() *gorm.DB {
	return db
}
