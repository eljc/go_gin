package main

import (
	"api-go-gin/database"
	"api-go-gin/models"
	"api-go-gin/routes"
)

func main() {
	database.ConnectDB()
	models.Students = []models.Student{
		{Id: 1, Name: "Camara", Course: "Golang"},
		{Id: 2, Name: "Aline", Course: "Java"},
	}
	routes.HandleRequests()
}
