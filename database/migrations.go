package database

import (
	"github.com/sanchitsamuel/webapp/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
