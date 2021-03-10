package database

import (
	"errors"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sato-takumi20-fixer/gin-training-api/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateInMemoryDbContext() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    if err != nil {
        panic(errors.New("Failed to Connect Database"))
    }
	Migrate(db)
    return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Formula{})
}