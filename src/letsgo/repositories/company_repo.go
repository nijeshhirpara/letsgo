package repositories

import (
	"letsgo/models"
	"log"

	"gorm.io/gorm"
)

// CompanyRepo implements models.CompanyRepository interface
type CompanyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) *CompanyRepo {
	return &CompanyRepo{
		db: db,
	}
}

// CreateCompany creats a company
func (cRepo *CompanyRepo) CreateCompany(c models.Company) error {
	res := cRepo.db.Create(&c)
	return res.Error
}

// ListCompanies retrives all companies
func (cRepo *CompanyRepo) ListCompanies() (companies []models.Company) {
	result := cRepo.db.Preload("Teams").Find(&companies)

	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	return
}

// FindCompanyByID finds a company by ID
func (cRepo *CompanyRepo) FindCompanyByID(companyID uint) (c models.Company, err error) {
	res := cRepo.db.First(&c, companyID)

	if res.Error != nil {
		err = res.Error
	}

	return
}

// AddTeamToCompany adds team to the company
func (cRepo *CompanyRepo) AddTeamToCompany(c models.Company, t models.Team) error {
	cRepo.db.Model(&c).Association("Teams").Append(&t)
	cRepo.db.Save(&c)
	return nil
}
