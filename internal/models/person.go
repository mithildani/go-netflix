package models

import "gorm.io/gorm"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Person struct {
	gorm.Model
	Name   string
	Gender Gender
}
