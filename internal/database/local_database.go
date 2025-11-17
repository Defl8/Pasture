package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Wraps gorm.DB for custom implemenation
type LocalDatabase struct {
	*gorm.DB
	Name string
}

func (ld *LocalDatabase) Connect() error {
	db, err := gorm.Open(sqlite.Open(ld.Name), &gorm.Config{})
	if err != nil {
		return err
	}
	ld.DB = db
	return nil
}
