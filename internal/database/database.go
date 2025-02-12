package database

import (
	"fmt"
	"os"
	"server/test-api/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func ConnectDb() (*gorm.DB, error) {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbTcp := "@tcp(127.0.0.1:3306)/"
	dsn := dbUser + ":" + dbPassword + dbTcp + dbName + "?charset=utf8&parseTime=True"
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("gorm Db connection ", err)
		return nil, err
	}

	errorMigration := gormDb.AutoMigrate(&models.User{})
	if errorMigration != nil {
		fmt.Println("gorm Db migration ", errorMigration)
		return nil, errorMigration
	}

	database = gormDb
	return gormDb, nil
}

func GetDb() *gorm.DB {
	if database == nil {
		panic("Database not connected")
	}
	return database
}
