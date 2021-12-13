package repositories

import (
	"letsgo/models"
	"log"

	"gorm.io/gorm"
)

// UserRepo implements models.UserRepository interface
type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// CreateUser creats an user
func (uRepo *UserRepo) CreateUser(company models.Company, user models.User) error {
	user.Company = company
	res := uRepo.db.Create(&user)
	return res.Error
}

// ListUsersByCompany retrives all users within a company
func (uRepo *UserRepo) ListUsersByCompany(companyID uint) (users []models.User) {
	result := uRepo.db.Preload("Teams").Where("company_id = ?", companyID).Find(&users)

	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	return
}

// FindUserByEmail retrives user by email
func (uRepo *UserRepo) FindUserByEmail(email string) (user models.User, err error) {
	result := uRepo.db.Where("email = ?", email).Find(&user)

	if result.Error != nil {
		log.Println(result.Error.Error())
		err = result.Error
	}

	return
}

// FindUserByID retrives user by ID
func (uRepo *UserRepo) FindUserByID(userID uint) (user models.User, err error) {
	result := uRepo.db.First(&user, userID)

	if result.Error != nil {
		log.Println(result.Error.Error())
		err = result.Error
	}

	return
}

// AddTeamsToUser adds teams to the user
func (uRepo *UserRepo) AddTeamsToUser(u models.User, teams []models.Team) error {
	uRepo.db.Model(&u).Association("Teams").Append(&teams)
	uRepo.db.Save(&u)
	return nil
}
