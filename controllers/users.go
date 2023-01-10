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

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cost := 8
	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), cost)
	if err != nil {
		panic(err)
	}
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

	if err := c.ShouldBindJSON(&newUsers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cost := 8
	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUsers.Password), cost)
	if err != nil {
		panic(err)
	}
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
		users     structs.Users
		jsonUsers structs.Users
		result    gin.H
	)
	if err := c.ShouldBindJSON(&jsonUsers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := idb.DB.Where("email = ?", jsonUsers.Email).First(&users).Error
	if err != nil {
		result = gin.H{
			"result": "Email not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(jsonUsers.Password))
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
