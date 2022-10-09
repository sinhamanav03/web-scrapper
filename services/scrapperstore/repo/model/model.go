package model

import "gorm.io/gorm"

type Item struct {
	Name        string
	ImageURL    string
	Description string
	Price       string
	Reviews     int
	gorm.Model
}
