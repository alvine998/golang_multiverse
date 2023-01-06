package structs

import "github.com/jinzhu/gorm"

type Profiles struct {
	gorm.Model
	Name    string
	Email   string
	Address string
	Phone   string
	Lat     string
	Long    string
}
