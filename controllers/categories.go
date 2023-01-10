package controllers

import (
	"api/webservice-multiverse/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to get data with id
func (idb *InDB) GetCategory(c *gin.Context) {

	var (
		categories structs.Categories
		result     gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id=?", id).First(&categories).Error
	if err != nil {
		result = gin.H{
			"result": err,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": categories,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Get All Data
func (idb *InDB) GetCategories(c *gin.Context) {
	search := c.Query("search")
	searchName := "%" + search + "%"
	var (
		categories []structs.Categories
		result     gin.H
	)

	idb.DB.Where("name LIKE ?", searchName).Find(&categories)
	if len(categories) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": categories,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Create user
func (idb *InDB) CreateCategory(c *gin.Context) {
	var (
		categories structs.Categories
		result     gin.H
	)

	if err := c.ShouldBindJSON(&categories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idb.DB.Create(&categories)
	result = gin.H{
		"result": categories,
	}
	c.JSON(http.StatusOK, result)
}

// Update user
func (idb *InDB) UpdateCategory(c *gin.Context) {
	id := c.Query("id")

	var (
		categories    structs.Categories
		newCategories structs.Categories
		result        gin.H
	)

	err := idb.DB.First(&categories, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}
	if err := c.ShouldBindJSON(&newCategories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = idb.DB.Model(&categories).Updates(newCategories).Error

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
func (idb *InDB) DeleteCategory(c *gin.Context) {
	var (
		categories structs.Categories
		result     gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&categories, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}
	err = idb.DB.Delete(&categories).Error
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
