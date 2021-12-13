package models

import "gorm.io/gorm"

// User table contains information about each user
type User struct {
	gorm.Model
	ID        uint
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CompanyID uint    `json:"-"`
	Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Teams     []Team  `gorm:"many2many:user_teams;"`
}

// UserRepository interface
type UserRepository interface {
	ListUsersByCompany(companyID uint) (users []User)
	CreateUser(company Company, u User) error
	FindUserByEmail(email string) (user User, err error)
	FindUserByID(userID uint) (user User, err error)
	AddTeamsToUser(u User, teams []Team) error
}
