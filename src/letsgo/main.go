package main

import (
	"letsgo/database"
	"letsgo/model"
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

	database.Connect(os.Getenv("DBPATH"))
	if err := database.DB.AutoMigrate(&model.Company{}, &model.Team{}, &model.User{}); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	router := route.RegisterRoutes(ctx)

	log.Fatal(http.ListenAndServe(host, handlers.LoggingHandler(os.Stdout, router)))
}
