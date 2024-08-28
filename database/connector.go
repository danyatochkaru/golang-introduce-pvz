package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(
		new(models.Order),
		new(models.Product),
		new(models.Status),
		new(models.Cart),
		new(models.CartProducts),
	)
	_ = db.SetupJoinTable(new(models.Cart), "Products", new(models.CartProducts))

	dbConnection = db

	return dbConnection
}
