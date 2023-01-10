package controllers

import (
	"api/webservice-multiverse/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to get data with id
func (idb *InDB) GetTestimony(c *gin.Context) {

	var (
		testimonies structs.Testimonies
		result      gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id=?", id).First(&testimonies).Error
	if err != nil {
		result = gin.H{
			"result": err,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": testimonies,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Get All Data
func (idb *InDB) GetTestimonies(c *gin.Context) {
	search := c.Query("search")
	searchName := "%" + search + "%"
	var (
		testimonies []structs.Testimonies
		result      gin.H
	)

	idb.DB.Where("name LIKE ?", searchName).Find(&testimonies)
	if len(testimonies) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": testimonies,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Create user
func (idb *InDB) CreateTestimony(c *gin.Context) {
	var (
		testimonies structs.Testimonies
		result      gin.H
	)

	name := c.PostForm("name")
	notes := c.PostForm("notes")

	testimonies.Name = name
	testimonies.Notes = notes

	idb.DB.Create(&testimonies)
	result = gin.H{
		"result": testimonies,
	}
	c.JSON(http.StatusOK, result)
}

// Update user
func (idb *InDB) UpdateTestimony(c *gin.Context) {
	id := c.Query("id")

	name := c.PostForm("name")
	notes := c.PostForm("notes")

	var (
		testimonies    structs.Testimonies
		newTestimonies structs.Testimonies
		result         gin.H
	)

	err := idb.DB.First(&testimonies, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}

	newTestimonies.Name = name
	newTestimonies.Notes = notes

	err = idb.DB.Model(&testimonies).Updates(newTestimonies).Error

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
func (idb *InDB) DeleteTestimony(c *gin.Context) {
	var (
		testimonies structs.Testimonies
		result      gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&testimonies, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}
	err = idb.DB.Delete(&testimonies).Error
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
