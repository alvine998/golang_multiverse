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

	if err := c.ShouldBindJSON(&profiles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idb.DB.Create(&profiles)
	result = gin.H{
		"result": profiles,
	}
	c.JSON(http.StatusOK, result)
}

// Update user
func (idb *InDB) UpdateProfile(c *gin.Context) {
	id := c.Query("id")

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

	if err := c.ShouldBindJSON(&newProfiles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
