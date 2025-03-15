package main

import (
	"github.com/amaan287/flightApiGo/initilizers"
	"github.com/amaan287/flightApiGo/models"
)

func init() {
	initilizers.LoadEnv()
	initilizers.ConnectToDB()
}
func main() {
	initilizers.DB.AutoMigrate(&models.User{})
}
