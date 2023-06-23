package models

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	MovieID       uint
	Audios        []Audio
	Videos        []Video
	TrailerVideos []Video
	TrailerAudios []Audio
	Thumbnail     string
}

type Audio struct {
	gorm.Model
	Language Language
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
