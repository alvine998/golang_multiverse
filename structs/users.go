package structs

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Username string
	Phone    string
}
