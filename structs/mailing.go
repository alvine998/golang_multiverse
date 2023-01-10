package structs

import "github.com/jinzhu/gorm"

type Mailing struct {
	gorm.Model
	Email   string
	Subject string
	Message string
}
