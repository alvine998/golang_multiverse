package structs

import "github.com/jinzhu/gorm"

type Subcategories struct {
	gorm.Model
	Name        string
	Category_id int
	Notes       string
}
