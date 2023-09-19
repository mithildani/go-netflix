package models

import "gorm.io/gorm"

type Role string

const (
	Director Role = "director"
	Actor    Role = "actor"
	Writer   Role = "writer"
	Producer Role = "producer"
)

type Crew struct {
	gorm.Model
	Movie  Movie
	Person Person
	Role   Role
}
