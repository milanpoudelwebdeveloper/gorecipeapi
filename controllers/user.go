package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

// GetProfile godoc
// @Summary Get user profile
// @Description Get user profile
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User	"User profile"
// @Router /user/profile [get]
func (ctrl UserController) UpdateProfile(c *gin.Context) {

}

// ChangePassword godoc
// @Summary Change password
// @Description Change password
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"Password changed"
// @Router /user/change/password [put]
func (ctrl UserController) ChangePassword(c *gin.Context) {

}
