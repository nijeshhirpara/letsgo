package helpers

import (
	"letsgo/database"
	"letsgo/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var TestC models.Company = models.Company{
	Name:        "Test Company 1",
	Description: "This is the Test Company 1",
}

var TestTm models.Team = models.Team{
	Name:        "Test Team 1",
	Description: "This is the Test Team 1",
}

func PrepareTestDB() *gorm.DB {
	// load .env file
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatalf("Error loading .env.test file")
	}

	// connect to the database
	db := database.Connect(os.Getenv("DBPATH"))
	if err := db.AutoMigrate(&models.Company{}, &models.Team{}, &models.User{}); err != nil {
		panic(err)
	}

	return db
}

func DropDBTables(db *gorm.DB) {
	db.Migrator().DropTable(&models.Company{})
	db.Migrator().DropTable(&models.Team{})
	db.Migrator().DropTable(&models.User{})
}
