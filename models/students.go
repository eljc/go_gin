package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Id     int64  `json:"id"`
	Name   string `json:"name" validate:"nonzero"`
	Course string `json:"course" validate:"nonzero"`
}

func ValidateDataStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}

var Students []Student
