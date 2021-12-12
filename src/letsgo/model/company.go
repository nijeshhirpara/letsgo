package model

import (
	"letsgo/database"
	"log"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID          uint
	Name        string `gorm:"unique"`
	Description string
	Teams       []Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func CreateCompany(c Company) error {
	res := database.DB.Create(&c)
	return res.Error
}

func ListCompanies() (companies []Company) {
	result := database.DB.Find(&companies)

	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	return
}
