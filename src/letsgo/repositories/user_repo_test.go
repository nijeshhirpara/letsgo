package repositories

import (
	"letsgo/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	comp, err := testCompanyRepo.FindCompanyByID(1)
	assert.Nil(t, err, "Error should be nil")

	err = testUserRepo.CreateUser(comp, helpers.TestUsr)
	assert.Nil(t, err, "Error should be nil")
}

func TestCreateDuplicateUser(t *testing.T) {
	comp, err := testCompanyRepo.FindCompanyByID(1)
	assert.Nil(t, err, "Error should be nil")

	err = testUserRepo.CreateUser(comp, helpers.TestUsr)
	assert.NotNil(t, err, "Error should be nil")
	assert.Equal(t, "UNIQUE constraint failed: users.email", err.Error())
}

func TestListUsersByCompany(t *testing.T) {
	users := testUserRepo.ListUsersByCompany(1)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, helpers.TestUsr.Name, users[0].Name)
}

func TestFindUserByEmail(t *testing.T) {
	usr, err := testUserRepo.FindUserByEmail(helpers.TestUsr.Email)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, helpers.TestUsr.Name, usr.Name)
}

func TestFindUserByEmailNotExists(t *testing.T) {
	usr, err := testUserRepo.FindUserByEmail("xyz")
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, uint(0), usr.ID)
}
