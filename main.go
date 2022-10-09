package main

import (
	"assignment-kedua/database"
	"assignment-kedua/routers"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db := database.StartDB()
	routers.StartServer(db).Run()
}