package model

import (
	"letsgo/database"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
)

var testC Company = Company{
	Name:        "Test Company 1",
	Description: "This is the Test Company 1",
}

func init() {
	PrepareTestDB()
}

func TestCreateCompany(t *testing.T) {
	err := CreateCompany(testC)
	assert.Nil(t, err, "Error should be nil")
}

func TestCreateCompanyWithError(t *testing.T) {
	err := CreateCompany(testC)
	assert.Equal(t, err.Error(), "UNIQUE constraint failed: companies.name")
}

func TestListCompanies(t *testing.T) {
	res := ListCompanies()
	assert.Equal(t, 1, len(res))
	assert.Equal(t, testC.Name, res[0].Name)
	assert.Equal(t, testC.Description, res[0].Description)

	DropDBTables()
}

func PrepareTestDB() {
	// load .env file
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatalf("Error loading .env.test file")
	}

	// connect to the database
	database.Connect(os.Getenv("DBPATH"))
	if err := database.DB.AutoMigrate(&Company{}, &Team{}, &User{}); err != nil {
		panic(err)
	}
}

func DropDBTables() {
	database.DB.Migrator().DropTable(&Company{})
	database.DB.Migrator().DropTable(&Team{})
	database.DB.Migrator().DropTable(&User{})
}
