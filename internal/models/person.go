package models

type Person struct {
	Name string `gorm:"primaryKey"`
	Gender string
}
