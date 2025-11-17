package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model        // This makes updated, deleted and created with ID automatically
	Title      string `gorm:"not null"`
	Content    string `gorm:"not null"`
	Published  bool   `gorm:"not null"`
}

func (p Post) TableName() string {
	return "posts"
}
