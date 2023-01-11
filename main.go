package main

import (
	"api/webservice-multiverse/config"
	"api/webservice-multiverse/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()
	router.Use(cors.New(CORSConfig()))

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

	// router profiles
	router.GET("profiles/:id", inDB.GetProfile)
	router.GET("profiles/", inDB.GetProfiles)
	router.POST("profiles/", inDB.CreateProfile)
	router.PATCH("profiles/", inDB.UpdateProfile)
	router.DELETE("profiles/:id", inDB.DeleteProfile)

	// router testimonies
	router.GET("testimonies/:id", inDB.GetTestimony)
	router.GET("testimonies/", inDB.GetTestimonies)
	router.POST("testimonies/", inDB.CreateTestimony)
	router.PATCH("testimonies/", inDB.UpdateTestimony)
	router.DELETE("testimonies/:id", inDB.DeleteTestimony)

	//router mailing
	router.POST("sending/mail", inDB.SendEmail)

	// Running Port
	router.Run("127.0.0.1:4000")
}
