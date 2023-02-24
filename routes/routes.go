package routes

import (
	"api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowStudents)
	r.POST("/student", controllers.New)
	r.GET("/student/:id", controllers.FindStudentByID)
	r.DELETE("/student/:id", controllers.DeleteById)
	r.PATCH("/student/:id", controllers.EditStudent)
	r.Run()
}
