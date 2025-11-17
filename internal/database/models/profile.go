package models

// this will just be the local user since the platform is self hosted
type Profile struct {
	Name      string `gorm:"not null"`
	Bio       string `gorm:"not null"`
	AvatarURL string
}
