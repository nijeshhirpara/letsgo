package model

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	ID          uint
	Name        string
	Description string
	CompanyID   uint
}
