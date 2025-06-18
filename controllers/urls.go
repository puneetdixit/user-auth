package controllers

import (
	"net/http"
	"user-auth/config"
	"user-auth/models"

	"github.com/gin-gonic/gin"
)

type URLInput struct {
	Url             string `json:"url" binding:"required,url"`
	Timeout         int    `json:"timeout" binding:"required,gte=1"`
	NotifyOnTimeout bool   `json:"notify_on_timeout" binding:"required"`
	NotifyEmail     string `json:"notify_email" binding:"required,email"`
}

type URLResponse struct {
	ID              uint   `json:"id"`
	Url             string `json:"url"`
	Timeout         int    `json:"timeout"`
	NotifyOnTimeout bool   `json:"notify_on_timeout"`
	NotifyEmail     string `json:"notify_email"`
}

func AddUrl(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser, _ := user.(*models.User)

	var input URLInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := models.URL{
		Url:             input.Url,
		Timeout:         input.Timeout,
		UserID:          currentUser.ID,
		NotifyOnTimeout: input.NotifyOnTimeout,
		NotifyEmail:     input.NotifyEmail,
	}
	config.DB.Create(&url)

	c.JSON(http.StatusCreated, gin.H{"message": "URL added successfully"})
}

func GetUrl(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser, _ := user.(*models.User)

	var urls []models.URL
	config.DB.Where("user_id = ?", currentUser.ID).Find(&urls)

	var response []URLResponse
	for _, u := range urls {
		response = append(response, URLResponse{
			ID:              u.ID,
			Url:             u.Url,
			Timeout:         u.Timeout,
			NotifyOnTimeout: u.NotifyOnTimeout,
			NotifyEmail:     u.NotifyEmail,
		})
	}

	c.JSON(http.StatusOK, gin.H{"urls": response})
}
