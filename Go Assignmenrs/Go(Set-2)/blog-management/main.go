package main

import (
	"blog-management/config"
	"blog-management/handlers"
	"blog-management/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Init()

	router := gin.Default()

	router.Use(middlewares.Logger())

	router.POST("/blog", handlers.CreateBlog)
	router.GET("/blog/:id", handlers.GetBlogByID)
	router.GET("/blogs", handlers.GetAllBlogs)
	router.PUT("/blog/:id", handlers.UpdateBlog)
	router.DELETE("/blog/:id", handlers.DeleteBlog)

	err := router.Run(":8090")
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
