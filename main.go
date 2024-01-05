package main

import (
	"net/http"

	"example.com/boilerplate/controllers"
	"example.com/boilerplate/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	models.ConnectDatabase()

	http.ListenAndServe(":8000", (handler))
}
