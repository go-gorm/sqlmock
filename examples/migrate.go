package main

import (
	"gorm.io/gorm"
)

// Migrate models
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Post{},
	)
}
