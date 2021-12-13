package repositories

import (
	"letsgo/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompany(t *testing.T) {
	err := testCompanyRepo.CreateCompany(helpers.TestC)
	assert.Nil(t, err, "Error should be nil")
}

func TestCreateCompanyWithError(t *testing.T) {
	err := testCompanyRepo.CreateCompany(helpers.TestC)
	assert.Equal(t, "UNIQUE constraint failed: companies.name", err.Error())
}

func TestListCompanies(t *testing.T) {
	res := testCompanyRepo.ListCompanies()
	assert.Equal(t, 1, len(res))
	assert.Equal(t, helpers.TestC.Name, res[0].Name)
	assert.Equal(t, helpers.TestC.Description, res[0].Description)
}

func TestFindCompanyByID(t *testing.T) {
	comp, err := testCompanyRepo.FindCompanyByID(1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, helpers.TestC.Name, comp.Name)
	assert.Equal(t, helpers.TestC.Description, comp.Description)
}

func TestFindCompanyByIDNotExist(t *testing.T) {
	_, err := testCompanyRepo.FindCompanyByID(10)
	assert.NotNil(t, err, "Error should not be nil")
	assert.Equal(t, "record not found", err.Error())
}

func TestAddTeamToCompany(t *testing.T) {
	comp, err := testCompanyRepo.FindCompanyByID(1)
	assert.Nil(t, err, "Error should be nil")

	err = testCompanyRepo.AddTeamToCompany(comp, helpers.TestTm)
	assert.Nil(t, err, "Error should be nil")

	res := testCompanyRepo.ListCompanies()
	assert.Equal(t, 1, len(res))
	assert.Equal(t, 1, len(res[0].Teams))
	assert.Equal(t, helpers.TestTm.Name, res[0].Teams[0].Name)
}
