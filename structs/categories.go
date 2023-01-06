package structs

import "github.com/jinzhu/gorm"

type Categories struct {
	gorm.Model
	Name  string
	Notes string
}
