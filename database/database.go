package database

import (
	"assignment-kedua/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

var (
	HOSTNAME = "localhost"
	USER     = "postgres"
	PASSWORD = "root"
	DBPORT   = "5432"
	DBNAME   = "learn_go"

	db *gorm.DB
	err error
)

func StartDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s, dbname=%s sslmode=disable", 
                HOSTNAME, USER, PASSWORD, DBPORT, DBNAME)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})

	return db
}
