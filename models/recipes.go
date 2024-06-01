package models

type Recipe struct {
	ID           uint   `json:"id" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Country      string `json:"country" binding:"required"`
	Instructions string `json:"instructions" binding:"required"`
}
