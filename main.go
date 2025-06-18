package main

import (
	"user-auth/config"
	"user-auth/controllers"
	"user-auth/middleware"
	"user-auth/models"
	"user-auth/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	utils.InitLogger()
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{}, &models.URL{})

	log.Info("Starting Server...")
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/user")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", controllers.Profile)

	urlRoutes := r.Group("/url")
	urlRoutes.Use(middleware.AuthMiddleware())
	urlRoutes.POST("/add", controllers.AddUrl)
	urlRoutes.GET("/get", controllers.GetUrl)
	r.Run(":8081")
}
