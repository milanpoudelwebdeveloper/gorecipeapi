package models

type Recipe struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Country      string `json:"country"`
	Instructions string `json:"instructions"`
}
