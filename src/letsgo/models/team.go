package models

import "gorm.io/gorm"

// Team table contains information about each team within the company
type Team struct {
	gorm.Model
	ID          uint
	Name        string `gorm:"index:idx_team,unique"`
	Description string
	CompanyID   uint `gorm:"index:idx_team,unique"`
}

// TeamRepository interface
type TeamRepository interface {
	ListTeamsByCompany(companyID uint) (teams []Team)
	FindCompanyTeamByName(companyID uint, name string) (t Team, err error)
	FindTeamByID(teamID uint) (t Team, err error)
}
