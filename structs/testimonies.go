package structs

import "github.com/jinzhu/gorm"

type Testimonies struct {
	gorm.Model
	Name  string
	Notes string
}
