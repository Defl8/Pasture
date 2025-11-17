package database

import (
	dbModels "github.com/Defl8/pasture/internal/database/models"
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

func (ld *LocalDatabase) AutoMigrate(models ...dbModels.Model) error {
	// The marker interface is not necessary, but looks nicer in the def
	dbItems := make([]any, len(models))

	for i, model := range models {
		dbItems[i] = model
	}

	err := ld.DB.AutoMigrate(dbItems...)
	if err != nil {
		return err
	}
	return nil
}

func (ld *LocalDatabase) Initialize() error {
	err := ld.AutoMigrate(&dbModels.Post{}, &dbModels.Profile{})
	if err != nil {
		return err
	}
	return nil
}

func (ld *LocalDatabase) Create(model dbModels.Model) error {
	return ld.DB.Create(model).Error
}

func (ld *LocalDatabase) Update(model dbModels.Model) error {
	return ld.DB.Save(model).Error
}

func (ld *LocalDatabase) Delete(model dbModels.Model) error {
	return ld.DB.Delete(model).Error
}

func (ld *LocalDatabase) GetPostByID(id uint) (*dbModels.Post, error) {
	var post dbModels.Post
	if err := ld.DB.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
