package controllers

import (
	"log"
	"net/http"
	"recipeapi/db"
	"recipeapi/models"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct{}

// GetCategories godoc
// @Summary get all categories
// @Schemes
// @Description get all categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} models.Categories
// @Router /categories [get]
func (ctrl *CategoriesController) GetCategories(c *gin.Context) {
	query := "SELECT * FROM category"
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while fetching categories",
		})
		return
	}
	defer rows.Close()
	var categories []models.Categories
	for rows.Next() {
		var category models.Categories
		err := rows.Scan(&category.ID, &category.Name, &category.CoverImage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while fetching categories",
			})
			return
		}
		categories = append(categories, category)

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "All categories",
		"data":    categories,
	})
}

// AddCategory godoc
// @Summary add a category
// @Schemes
// @Description add a category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Categories true "Category object that needs to be added"
// @Success 201 {string} string "Category created successfully"
// @Router /recipes [post]
func (ctrl *CategoriesController) AddCategory(c *gin.Context) {
	query := "INSERT INTO category (name, coverimage) VALUES ($1, $2)"
	var newCategory models.Categories
	err := c.BindJSON(&newCategory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error parsing body",
		})
		return
	}
	_, err = db.DB.Exec(query, newCategory.Name, newCategory.CoverImage)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong while adding a category",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
		"data":    newCategory,
	})

}
