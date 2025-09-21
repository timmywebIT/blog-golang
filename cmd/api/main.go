package main

import (
	"database/sql"
	"log"
	_ "rest-api-in-gin/docs"
	"rest-api-in-gin/internal/database"
	"rest-api-in-gin/internal/env"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// @title Go Gin Rest API
// @version 1.0
// @description A rest API in Go using Gin framework.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

// Apply the security definition to your endpoints
// @security BearerAuth

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {

	db, error := sql.Open("sqlite3", "./data.db")
	if error != nil {
		log.Fatal(error)
	}

	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "secret"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}

}
