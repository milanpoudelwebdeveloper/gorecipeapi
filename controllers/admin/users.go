package admin

import (
	"log"
	"net/http"
	"recipeapi/db"
	"recipeapi/models"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

type NewUser struct {
	Email string `json:"email" binding:"required"`
	Role  string `json:"role" binding:"required"`
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags admin
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User	"List of users"
// @Router /admin/users [get]
func (ctrl *AdminController) GetUsers(c *gin.Context) {
	query := "SELECT name, email, role, verified, profileimage FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error while fetching users",
		})
		return
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Name, &user.Email, &user.Role, &user.Verified, &user.ProfileImage)
		if err != nil {
			log.Fatal("Here is an error:", err)
			c.JSON(500, gin.H{

				"message": "Error while fetching users",
			})
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "All users",
		"data":    users,
	})

}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags admin
// @Accept  json
// @Produce  json
// @Param body body NewUser true "New User"
// @Success 201 {string} string	"New user created successfully"
// @Router /admin/users [post]
func (ctrl *AdminController) CreateUser(c *gin.Context) {
	var newUser NewUser
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad body request",
		})
		return
	}
	query := "INSERT into users (email, role) VALUES ($1, $2)"
	if _, err := db.DB.Query(query, newUser.Email, newUser.Role); err != nil {
		c.JSON(400, gin.H{
			"message": "Something went wrong while creating a user",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin: New user created successfully",
	})
}
