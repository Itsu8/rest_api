package main

import (
	"github.com/Itsu8/rest_api/initializers"
	"github.com/Itsu8/rest_api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	
}
