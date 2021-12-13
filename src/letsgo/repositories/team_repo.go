package repositories

import (
	"letsgo/models"
	"log"

	"gorm.io/gorm"
)

// TeamRepo implements models.TeamRepository interface
type TeamRepo struct {
	db *gorm.DB
}

func NewTeamRepo(db *gorm.DB) *TeamRepo {
	return &TeamRepo{
		db: db,
	}
}

// ListTeamsByCompany retrives all teams within a company
func (tRepo *TeamRepo) ListTeamsByCompany(companyID uint) (teams []models.Team) {
	result := tRepo.db.Where("company_id = ?", companyID).Find(&teams)

	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	return
}

// FindCompanyTeamByName retrives a team with a given name within a company (here, Team name is unique within a company)
func (tRepo *TeamRepo) FindCompanyTeamByName(companyID uint, name string) (t models.Team, err error) {
	err = nil
	result := tRepo.db.Where("name = ? AND company_id = ?", name, companyID).Find(&t)

	if result.Error != nil {
		log.Println(result.Error.Error())
		err = result.Error
	}

	log.Println("result", t)

	return
}
