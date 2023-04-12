package config

import (
	"api/webservice-multiverse/structs"
	"os"

	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// Connection to DB
func DBInit() *gorm.DB {
	errors := godotenv.Load()
	if errors != nil {
		fmt.Println(errors)
	}

	Username := os.Getenv("DB_USERNAME")
	Name := os.Getenv("DB_NAME")
	Password := os.Getenv("DB_PASSWORD")
	if Username != "" {
		Username = "root"
	}
	if Name != "" {
		Name = "multiverse_go"
	}
	Url := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?parseTime=true", Username, Name, Password)

	db, err := gorm.Open("mysql", Url)
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(structs.Users{})
	return db
}
