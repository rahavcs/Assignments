package main

import (
	"ecom-inventory/handlers"
	"ecom-inventory/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the E-commerce Inventory Management API!",
		})
	})

	// Middleware
	router.Use(middleware.Logger())

	// Product Routes
	router.POST("/product", handlers.CreateProduct)
	router.GET("/product/:id", handlers.GetProductByID)
	router.PUT("/product/:id", handlers.UpdateProductStock)
	router.DELETE("/product/:id", handlers.DeleteProduct)

	// Start the server
	err := router.Run(":8090")
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
