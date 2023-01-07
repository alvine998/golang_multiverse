package structs

import "github.com/jinzhu/gorm"

type Products struct {
	gorm.Model
	Name        string
	Stock       int
	Category_id int
	Price       int
	Notes       string
	Status      int
}
