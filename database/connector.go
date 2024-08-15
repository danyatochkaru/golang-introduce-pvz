package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pvz/models"
)

var dbConnection *gorm.DB

func GetDBConnection() *gorm.DB {
	if dbConnection == nil {
		connectDB()
	}

	return dbConnection
}

func connectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(&models.Order{})

	dbConnection = db

	return dbConnection
}
