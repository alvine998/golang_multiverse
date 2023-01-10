package controllers

import (
	"api/webservice-multiverse/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to get data with id
func (idb *InDB) GetProfile(c *gin.Context) {

	var (
		profiles structs.Profiles
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id=?", id).First(&profiles).Error
	if err != nil {
		result = gin.H{
			"result": err,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": profiles,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Get All Data
func (idb *InDB) GetProfiles(c *gin.Context) {
	search := c.Query("search")
	searchName := "%" + search + "%"
	var (
		profiles []structs.Profiles
		result   gin.H
	)

	idb.DB.Where("name LIKE ?", searchName).Find(&profiles)
	if len(profiles) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": profiles,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Create user
func (idb *InDB) CreateProfile(c *gin.Context) {
	var (
		profiles structs.Profiles
		result   gin.H
	)

	name := c.PostForm("name")
	email := c.PostForm("email")
	address := c.PostForm("address")
	phone := c.PostForm("phone")
	lat := c.PostForm("lat")
	long := c.PostForm("long")

	profiles.Name = name
	profiles.Email = email
	profiles.Address = address
	profiles.Phone = phone
	profiles.Lat = lat
	profiles.Long = long

	idb.DB.Create(&profiles)
	result = gin.H{
		"result": profiles,
	}
	c.JSON(http.StatusOK, result)
}

// Update user
func (idb *InDB) UpdateProfile(c *gin.Context) {
	id := c.Query("id")

	name := c.PostForm("name")
	email := c.PostForm("email")
	address := c.PostForm("address")
	phone := c.PostForm("phone")
	lat := c.PostForm("lat")
	long := c.PostForm("long")

	var (
		profiles    structs.Profiles
		newProfiles structs.Profiles
		result      gin.H
	)

	err := idb.DB.First(&profiles, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}

	newProfiles.Name = name
	newProfiles.Email = email
	newProfiles.Address = address
	newProfiles.Phone = phone
	newProfiles.Lat = lat
	newProfiles.Long = long

	err = idb.DB.Model(&profiles).Updates(newProfiles).Error

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
func (idb *InDB) DeleteProfile(c *gin.Context) {
	var (
		profiles structs.Profiles
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&profiles, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}
	err = idb.DB.Delete(&profiles).Error
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
