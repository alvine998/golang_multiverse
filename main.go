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

	// Router users
	router.GET("users/:id", inDB.GetUser)
	router.GET("users/", inDB.GetUsers)
	router.POST("users/", inDB.CreateUser)
	router.POST("users/auth", inDB.AuthUser)
	router.PATCH("users/", inDB.UpdateUser)
	router.DELETE("users/:id", inDB.DeleteUser)

	// router products
	router.GET("products/:id", inDB.GetProduct)
	router.GET("products/", inDB.GetProducts)
	router.POST("products/", inDB.CreateProduct)
	router.PATCH("products/", inDB.UpdateProduct)
	router.DELETE("products/:id", inDB.DeleteProduct)

	// router categories
	router.GET("categories/:id", inDB.GetCategory)
	router.GET("categories/", inDB.GetCategories)
	router.POST("categories/", inDB.CreateCategory)
	router.PATCH("categories/", inDB.UpdateCategory)
	router.DELETE("categories/:id", inDB.DeleteCategory)

	// router subcategories
	router.GET("subcategories/:id", inDB.GetSubcategory)
	router.GET("subcategories/", inDB.GetSubcategories)
	router.POST("subcategories/", inDB.CreateSubcategory)
	router.PATCH("subcategories/", inDB.UpdateSubcategory)
	router.DELETE("subcategories/:id", inDB.DeleteSubcategory)

	// router contents
	router.GET("contents/:id", inDB.GetContent)
	router.GET("contents/", inDB.GetContents)
	router.POST("contents/", inDB.CreateContent)
	router.PATCH("contents/", inDB.UpdateContent)
	router.DELETE("contents/:id", inDB.DeleteContent)

	// Running Port
	router.Run("127.0.0.1:4000")
}
