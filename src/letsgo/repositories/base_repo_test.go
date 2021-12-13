package repositories

import (
	"letsgo/helpers"
	"os"
	"testing"
)

var testCompanyRepo *CompanyRepo
var testTeamRepo *TeamRepo

func TestMain(m *testing.M) {
	db := helpers.PrepareTestDB()

	testCompanyRepo = NewCompanyRepo(db)
	testTeamRepo = NewTeamRepo(db)

	code := m.Run()

	helpers.DropDBTables(db)
	os.Exit(code)
}
