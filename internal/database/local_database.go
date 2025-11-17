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

func NewLocalDatabase(name string) (*LocalDatabase, error) {
	ld := &LocalDatabase{Name: name}

	// Initialize the connection on struct creation
	if err := ld.Connect(); err != nil {
		return nil, err
	}

	return ld, nil
}

func (ld *LocalDatabase) Close() error {
	sqlDB, err := ld.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}


func (ld *LocalDatabase) AutoMigrate(models ...any) {
	ld.DB.AutoMigrate(models...)
}
