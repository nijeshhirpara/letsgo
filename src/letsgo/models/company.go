package models

import "gorm.io/gorm"

// Company table contains information about each company
type Company struct {
	gorm.Model
	ID          uint
	Name        string `gorm:"unique"`
	Description string
	Teams       []Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// CompanyRepository interface
type CompanyRepository interface {
	CreateCompany(c Company) error
	ListCompanies() (companies []Company)
	FindCompanyByID(companyID uint) (c Company, err error)
	AddTeamToCompany(c Company, t Team) error
}
