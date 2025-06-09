package controllers

import (
	"user-auth/models"
	"user-auth/config"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UrlInput struct {
	Url             string `json:"url" binding:"required,url"`
	Timeout         int    `json:"timeout" binding:"required,gte=1"`
	NotifyOnTimeout bool   `json:"notify_on_timeout" binding:"required"`
	NotifyEmail     string `json:"notify_email" binding:"required,email"`
}


func AddUrl(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser, _ := user.(*models.User)

	var input UrlInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := models.URL{
		Url: input.Url, 
		Timeout: input.Timeout, 
		UserID: currentUser.ID, 
		NotifyOnTimeout: input.NotifyOnTimeout, 
		NotifyEmail: input.NotifyEmail,
	}
	config.DB.Create(&url)
	
	c.JSON(http.StatusCreated, gin.H{"message": "URL added successfully"})
}