package structs

import "github.com/jinzhu/gorm"

type Contents struct {
	gorm.Model
	Title string
	Image string
	Notes string
}
