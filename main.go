package main

import (
	"api/webservice-multiverse/config"
	"api/webservice-multiverse/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	// Router
	router.GET("users/:id", inDB.GetUser)
	router.GET("users/", inDB.GetUsers)
	router.POST("users/", inDB.CreateUser)
	router.PATCH("users/", inDB.UpdateUser)
	router.DELETE("users/:id", inDB.DeleteUser)

	// Running Port
	router.Run("127.0.0.1:4000")
}
