package main

import (
	"user-auth/controllers"
	"user-auth/middleware"
	"user-auth/models"
	"user-auth/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	utils.InitLogger()
	models.ConnectDatabase()

	log.Info("Stating Server...")
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/user")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", controllers.Profile)

	r.Run(":8080")
}
