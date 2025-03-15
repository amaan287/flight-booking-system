package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/amaan287/flightApiGo/controller"
	"github.com/amaan287/flightApiGo/initilizers"
	"github.com/gin-gonic/gin"
)

func init() {
	initilizers.LoadEnv()
	initilizers.ConnectToDB()
}

func main() {
	port := os.Getenv("PORT")
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello from server")
	})
	router.POST("/Signup", controller.Signup)
	fmt.Printf("Server is running on http://localhost%s", port)
	http.ListenAndServe(port, router)
}
