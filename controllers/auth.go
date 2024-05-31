package controllers

import (
	"database/sql"
	"net/http"
	"recipeapi/db"
	"recipeapi/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Country  string `json:"country" binding:"required"`
}

type AuthController struct{}

func (ctrl *AuthController) Login(c *gin.Context) {
	var loginRequest LoginRequest

	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"message": "Couldn't parse login body", "error": err.Error()})
		return
	}
	parsedEmail := loginRequest.Email
	parsedPassword := loginRequest.Password
	query := "SELECT email, password FROM users WHERE email = $1"
	var user models.User
	err := db.DB.QueryRow(query, parsedEmail).Scan(&user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "User with that email not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while fetching user",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(parsedPassword))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect password"})
		return
	}
	c.JSON(200, gin.H{"message": "Login successful"})
}

func (ctrl AuthController) Register(c *gin.Context) {
	var registerRequest RegisterRequest
	if err := c.BindJSON(&registerRequest); err != nil {
		c.JSON(400, gin.H{"message": "Couldn't parse register body", "error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while hashing password",
		})
		return
	}
	query := "INSERT INTO users (name, email, password, country) VALUES ($1, $2, $3, $4)"
	_, err = db.DB.Exec(query, registerRequest.Name, registerRequest.Email, hashedPassword, registerRequest.Country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating user",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}
