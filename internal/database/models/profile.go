package models

import "gorm.io/gorm"

// this will just be the local user since the platform is self hosted
type Profile struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Bio       string `gorm:"not null"`
	AvatarURL string
}

func (Profile) TableName() string {
	return "profile"
}
