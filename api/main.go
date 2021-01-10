package main

import (
	"log"
	"github.com/gin-gonic/gin"

	controllers "go_test/controllers"

)


func main() {
	// Set up HTTP server
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Next()
	})

	r.GET("/deliveries/invalid_combination", controllers.FindInvalidDeliveries)

	// Run HTTP server
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
