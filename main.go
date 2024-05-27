package main

import (
	"fmt"
	"log"
	"portal-data-kalsel/database"
	"portal-data-kalsel/model"
	"portal-data-kalsel/routes"
)

// @title Service Portal-Data-Kalsel
// @description Portal-Data-Kalsel
// @host localhost:8080
// @BasePath /
func main() {
	database.Connect()
	if err := database.DB.AutoMigrate(&model.Artikel{}, &model.Infografis{}); err != nil {
		log.Println("Failed to migrate database:", err)
		fmt.Println("===================================")
	}
	routes.Routes()
}
