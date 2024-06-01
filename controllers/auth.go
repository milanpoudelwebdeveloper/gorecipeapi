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

// Login godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body LoginRequest true "Login Request"
// @Success 200 {string} string	"Login successful"
// @Router /auth/login [post]
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

// Register godoc
// @Summary Register
// @Description Register
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body RegisterRequest true "Register Request"
// @Success 201 {string} string	"User created successfully"
// @Router /auth/register [post]
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

// VerifyAccount godoc
// @Summary Verify account
// @Description Verify account
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Email verified"
// @Router /auth/verify [put]
func (ctrl AuthController) VerifyAccount(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Email verified",
	})
}

// ForgotPassword godoc
// @Summary Forgot password
// @Description Forgot password
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Forgot password"
// @Router /auth/forgot/password [post]
func (ctrl AuthController) ForgotPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Forgot password",
	})
}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset password
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Reset password"
// @Router /auth/reset/password [put]
func (ctrl AuthController) ResetPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Reset password",
	})
}

// ResendVerification godoc
// @Summary Resend verification
// @Description Resend verification
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Resend verification"
// @Router /auth/resend/verification [get]
func (ctrl AuthController) ResendVerification(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Resend verification",
	})
}

// Logout godoc
// @Summary Logout
// @Description Logout
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Logout successful"
// @Router /auth/logout [get]
func (ctrl AuthController) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logout successful",
	})
}
