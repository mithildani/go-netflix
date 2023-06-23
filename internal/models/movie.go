package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title    string
	Year     uint16
	Summary  string
	Duration string
	Media    Media
	Genres   []Genre `gorm:"many2many:movie_genres;"`
	Director []Person
	Cast     []Person
	Writer   []Person
}
