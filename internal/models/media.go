package models

import "gorm.io/gorm"

type ContentType string

const (
	Default ContentType = "default"
	Trailer ContentType = "trailer"
)

type Audio struct {
	gorm.Model
	Movie    Movie
	Type     ContentType
	Language Language
	URL      string
}

type Language struct {
	gorm.Model
	ShortName   string
	DisplayName string
}

type Video struct {
	gorm.Model
	Movie   Movie
	Type    ContentType
	Quality Quality
	URL     string
}

type Quality struct {
	gorm.Model
	Level uint8
	Name  string
}
