package controllers

import (
	"api/webservice-multiverse/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to get data with id
func (idb *InDB) GetSubcategory(c *gin.Context) {

	var (
		subcategories structs.Subcategories
		result        gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id=?", id).First(&subcategories).Error
	if err != nil {
		result = gin.H{
			"result": err,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": subcategories,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Get All Data
func (idb *InDB) GetSubcategories(c *gin.Context) {
	search := c.Query("search")
	searchName := "%" + search + "%"
	var (
		subcategories []structs.Subcategories
		result        gin.H
	)

	idb.DB.Where("name LIKE ?", searchName).Find(&subcategories)
	if len(subcategories) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": subcategories,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Create user
func (idb *InDB) CreateSubcategory(c *gin.Context) {
	var (
		subcategories structs.Subcategories
		result        gin.H
	)

	name := c.PostForm("name")
	notes := c.PostForm("notes")

	subcategories.Name = name
	subcategories.Notes = notes

	idb.DB.Create(&subcategories)
	result = gin.H{
		"result": subcategories,
	}
	c.JSON(http.StatusOK, result)
}

// Update user
func (idb *InDB) UpdateSubcategory(c *gin.Context) {
	id := c.Query("id")

	name := c.PostForm("name")
	notes := c.PostForm("notes")

	var (
		subcategories    structs.Subcategories
		newSubcategories structs.Subcategories
		result           gin.H
	)

	err := idb.DB.First(&subcategories, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}

	newSubcategories.Name = name
	newSubcategories.Notes = notes

	err = idb.DB.Model(&subcategories).Updates(newSubcategories).Error

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
func (idb *InDB) DeleteSubcategory(c *gin.Context) {
	var (
		subcategories structs.Subcategories
		result        gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&subcategories, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}
	err = idb.DB.Delete(&subcategories).Error
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
