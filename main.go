package main

import (
	"user-auth/controllers"
	"user-auth/middleware"
	"user-auth/utils"
	"user-auth/config"

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

	log.Info("Starting Server...")
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/user")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", controllers.Profile)

	r.Run(":8081")
}
