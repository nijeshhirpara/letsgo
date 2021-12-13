package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect opens a connection to the database
func Connect(dbpath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
