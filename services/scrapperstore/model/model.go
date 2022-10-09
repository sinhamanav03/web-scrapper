package model

type Item struct {
	Name        string `json:"name"`
	ImageURL    string `json:"imageURL"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Reviews     int    `json:"totalReviews"`
}
