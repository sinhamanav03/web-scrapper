package model

type (
	ScrapeUrlResponse struct {
		Name        string `json:"name"`
		ImageURL    string `json:"imageURL"`
		Description string `json:"description"`
		Price       string `json:"price"`
		Reviews     string `json:"reviews"`
	}
)
