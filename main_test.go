package main

import (
	"api-go-gin/controllers"
	"api-go-gin/database"
	"api-go-gin/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetUpRoutesTests() *gin.Engine {
	routes := gin.Default()
	return routes
}

func MockStudent() {
	student := models.Student{Name: "Student Name", Course: "Golang"}
	database.DB.Create(&student)
	ID = int(student.Id)
}

func DeletaStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestStatusCode(t *testing.T) {
	r := SetUpRoutesTests()
	r.GET("/students", controllers.ShowStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditStudent(t *testing.T) {
	database.ConnectDB()
	MockStudent()
	defer DeletaStudentMock()
	r := SetUpRoutesTests()
	r.PATCH("/student/:id", controllers.EditStudent)
	student := models.Student{Name: "Name of Student2", Course: "Course2"}
	valorJson, _ := json.Marshal(student)
	pathEdit := "/student/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathEdit, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var studentMockUpdated models.Student
	json.Unmarshal(resposta.Body.Bytes(), &studentMockUpdated)
	assert.Equal(t, "Name of Student2", studentMockUpdated.Name)
	assert.Equal(t, "Course2", studentMockUpdated.Course)

}

func TestDeletStudent(t *testing.T) {
	database.ConnectDB()
	MockStudent()
	defer DeletaStudentMock()
	r := SetUpRoutesTests()
	r.DELETE("/student/:id", controllers.DeleteById)
	pathDelete := "/student/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDelete, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)

}
