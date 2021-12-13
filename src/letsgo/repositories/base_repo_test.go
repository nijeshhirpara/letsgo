package repositories

import (
	"letsgo/helpers"
	"os"
	"testing"
)

var testCompanyRepo *CompanyRepo
var testTeamRepo *TeamRepo
var testUserRepo *UserRepo

func TestMain(m *testing.M) {
	db := helpers.PrepareTestDB()

	testCompanyRepo = NewCompanyRepo(db)
	testTeamRepo = NewTeamRepo(db)
	testUserRepo = NewUserRepo(db)

	code := m.Run()

	helpers.DropDBTables(db)
	os.Exit(code)
}
