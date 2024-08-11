package main

import (
	"log"
	"os"

	"fholl.net/go-pg-sample/database"
	"fholl.net/go-pg-sample/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"fholl.net/go-pg-sample/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	cfg := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DB:       os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db, err := database.NewConnection(&cfg)
	if err != nil {
		log.Fatal("Could not load the database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Booking{}, &models.Client{}, &models.Contractor{}, &models.Contact{})

	r := routes.Routes{DB: db.Debug()}
	r.Setup(e)

	v := routes.Views{}
	v.Setup(e)

	e.Logger.Fatal(e.Start(":8080"))
}
