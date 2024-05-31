package controllers

import (
	"fmt"
	"log"
	"net/http"
	"recipeapi/db"
	"recipeapi/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

// GetUsers godoc
// @Summary get all users
// @Schemes
// @Description get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /users [get]
func (ctrl UserController) GetUsers(c *gin.Context) {
	query := "SELECT * FROM users ORDER BY id"
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while fetching users",
		})
		return
	}
	fmt.Println("the rows are", rows)
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Country, &user.Verified, &user.Role, &user.ProfileImage)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{
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

// AddUser godoc
// @Summary Post a user
// @Schemes
// @Description Post a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object that needs to be added"
// @Success 201 {string} string "User added successfully"
// @Router /users [post]
func (ctrl UserController) AddUser(c *gin.Context) {
	query := "INSERT INTO users (name, email, password, country, verified, role, profileimage) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *"
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payload request",
		})
		return
	}
	data, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while hashing password",
		})
		return

	}
	newUser.Password = string(data)
	_, err = db.DB.Exec(query, newUser.Name, newUser.Email, newUser.Password, newUser.Country, newUser.Verified, newUser.Role, newUser.ProfileImage)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating user",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    newUser,
	})
}
