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
	result := uRepo.db.Where("company_id = ?", companyID).Find(&users)

	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	return
}

// FindUserByEmail retrives user by email
func (uRepo *UserRepo) FindUserByEmail(email string) (users models.User, err error) {
	result := uRepo.db.Where("email = ?", email).Find(&users)

	if result.Error != nil {
		log.Println(result.Error.Error())
		err = result.Error
	}

	return
}
