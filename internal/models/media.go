package models

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	MovieID       uint
	Audios        []Audio
	Videos        []Video
	TrailerAudios []Audio
	TrailerVideos []Video
	Thumbnail     string
}

type Audio struct {
	gorm.Model
	Language string `gorm:"unique"`
	URL      string
}

type Video struct {
	gorm.Model
	Quality Quality
	URL     string
}

type Quality struct {
	gorm.Model
	Level uint8
	Name  string
}
