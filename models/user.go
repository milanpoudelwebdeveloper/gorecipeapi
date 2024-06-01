package models

type User struct {
	ID           uint   `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Country      string `json:"country" binding:"required"`
	Role         string `json:"role" binding:"required"`
	ProfileImage string `json:"profile_image" binding:"required"`
	Verified     bool   `json:"verified" binding:"required"`
}
