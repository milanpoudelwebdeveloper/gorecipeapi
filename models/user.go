package models

type User struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Country      string `json:"country"`
	Role         string `json:"role"`
	ProfileImage string `json:"profile_image"`
	Verified     bool   `json:"verified"`
}
