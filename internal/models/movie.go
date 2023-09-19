package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title     string
	Year      uint16
	Summary   string
	Thumbnail string

	Audios []Audio
	Videos []Video
	Genres []Genre  `gorm:"many2many:movie_genres;"`
	Crew   []Person `gorm:"many2many:crew;"`
}
