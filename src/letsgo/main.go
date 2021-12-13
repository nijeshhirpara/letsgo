package main

import (
	"letsgo/controller/api"
	"letsgo/database"
	"letsgo/models"
	"letsgo/repositories"
	"letsgo/route"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
)

func main() {
	log.Println("Let's Go")

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// connect to the database
	db := database.Connect(os.Getenv("DBPATH"))
	if err := db.AutoMigrate(&models.Company{}, &models.Team{}, &models.User{}); err != nil {
		panic(err)
	}

	// Create repos
	companyRepo := repositories.NewCompanyRepo(db)
	teamRepo := repositories.NewTeamRepo(db)
	userRepo := repositories.NewUserRepo(db)

	// Initiate controllers
	h := api.NewBaseHandler(companyRepo, teamRepo, userRepo)

	// Initiate router
	r := route.NewRouter(h)

	// create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// register routes
	router := r.RegisterRoutes(ctx)

	// start server
	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(host, handlers.LoggingHandler(os.Stdout, router)))
}
