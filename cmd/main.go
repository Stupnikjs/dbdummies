package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Stupnikjs/dbdummies/api"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	app := api.Application{
		Port: 8080,
	}

	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())

}
