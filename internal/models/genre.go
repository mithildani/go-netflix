package models

type Genre struct {
	Name   string  `gorm:"primaryKey"`
	Movies []Movie `gorm:"many2many:movie_genres"`
}
