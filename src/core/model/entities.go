package model

import "time"

// '{"title": "some_title", "text": "some_text", "url": "some_url"}'

type ContentDTO struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}

type BannerContent struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Text      string
	Url       string
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	IsActive  bool
}

type Feature struct {
	ID        uint `gorm:"primary_key"`
	ContentID uint `gorm:"not null"`
}

type Tag struct {
	ID        uint `gorm:"primary_key"`
	ContentID uint `gorm:"not null"`
}
