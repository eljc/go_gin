package database

import (
	"api-go-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	strConnection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConnection))
	if err != nil {
		log.Panic("Error to connect DB")
	}
	DB.AutoMigrate(&models.Student{})
}
