package structs

import "github.com/jinzhu/gorm"

type Products struct {
	gorm.Model
	Name        string
	Stock       int
	Category_id string
	Price       int
	Notes       string
	Status      int
}
