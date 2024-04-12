package model

import "time"

type ContentDTO struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}

type BannerContent struct {
	ID        int32 `gorm:"primary_key"`
	Title     string
	Text      string
	Url       string
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	IsActive  bool
}

type Feature struct {
	ID        int32 `gorm:"primary_key"`
	ContentID int32 `gorm:"not null"`
}

type Tag struct {
	ID        int32 `gorm:"primary_key"`
	ContentID int32 `gorm:"not null"`
}
