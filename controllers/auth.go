package controllers

import (
	"net/http"
	"user-auth/config"
	"user-auth/models"
	"user-auth/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	var userExists models.User
	if config.DB.Where("email = ? or username = ?", input.Email, input.Username).First(&userExists).RowsAffected != 0 {
		if userExists.Email == input.Email {
			c.JSON(http.StatusConflict, gin.H{"error": "User already registered with this email"})
		} else {
			c.JSON(http.StatusConflict, gin.H{"error": "User already registered with this username"})
		}
		return
	}

	config.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProfileResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, _ := utils.GenerateJWT(user)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Profile(c *gin.Context) {
	user, _ := c.Get("user")
	userData := user.(*models.User)

	log.Info("user", userData)
	profile := ProfileResponse{
		ID:       userData.ID,
		Email:    userData.Email,
		Username: userData.Username,
	}

	c.JSON(http.StatusOK, profile)
}
