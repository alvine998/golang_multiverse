package controllers

import (
	"api/webservice-multiverse/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// to get data with id
func (idb *InDB) GetUser(c *gin.Context) {

	var (
		users  structs.Users
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id=?", id).First(&users).Error
	if err != nil {
		result = gin.H{
			"result": err,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": users,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Get All Data users
func (idb *InDB) GetUsers(c *gin.Context) {
	search := c.Query("search")
	searchName := "%" + search + "%"
	var (
		users  []structs.Users
		result gin.H
	)

	idb.DB.Where("name LIKE ?", searchName).Find(&users)
	if len(users) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": users,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Create user
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		users  structs.Users
		result gin.H
	)

	name := c.PostForm("name")
	email := c.PostForm("email")
	username := c.PostForm("username")
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	cost := 8
	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		panic(err)
	}

	users.Name = name
	users.Email = email
	users.Username = username
	users.Phone = phone
	users.Password = string(hashedPassword)

	idb.DB.Create(&users)
	result = gin.H{
		"result": users,
	}
	c.JSON(http.StatusOK, result)
}

// Update user
func (idb *InDB) UpdateUser(c *gin.Context) {
	id := c.Query("id")

	name := c.PostForm("name")
	email := c.PostForm("email")
	username := c.PostForm("username")
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	var (
		users    structs.Users
		newUsers structs.Users
		result   gin.H
	)

	err := idb.DB.First(&users, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}

	cost := 8
	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		panic(err)
	}

	newUsers.Name = name
	newUsers.Email = email
	newUsers.Username = username
	newUsers.Phone = phone
	newUsers.Password = string(hashedPassword)

	err = idb.DB.Model(&users).Updates(newUsers).Error

	if err != nil {
		result = gin.H{
			"result": "update failed",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	} else {
		result = gin.H{
			"result": "successfully update data",
		}
		c.JSON(http.StatusOK, result)
	}
}

// Delete user by id
func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		users  structs.Users
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&users, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	err = idb.DB.Delete(&users).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
		c.JSON(http.StatusOK, result)
	}

}

// Login
func (idb *InDB) AuthUser(c *gin.Context) {
	var (
		users  structs.Users
		result gin.H
	)
	email := c.PostForm("email")
	password := c.PostForm("password")

	err := idb.DB.Where("email = ?", email).First(&users).Error
	if err != nil {
		result = gin.H{
			"result": "Email not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
		if err != nil {
			result = gin.H{
				"status": "Error",
				"result": "Invalid Password!",
				"Code":   http.StatusBadRequest,
			}
			c.JSON(http.StatusBadRequest, result)
		} else {
			result = gin.H{
				"status": "Success",
				"result": "Login Success",
				"Code":   http.StatusOK,
			}
			c.JSON(http.StatusOK, result)
		}
	}
}
