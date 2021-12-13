package models

import "gorm.io/gorm"

// User table contains information about each user
type User struct {
	gorm.Model
	ID        uint
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CompanyID uint
	Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Teams     []Team  `gorm:"many2many:user_teams;"`
}
