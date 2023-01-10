package controllers

import (
	"api/webservice-multiverse/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// to get data with id
func (idb *InDB) GetProduct(c *gin.Context) {

	var (
		products structs.Products
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id=?", id).First(&products).Error
	if err != nil {
		result = gin.H{
			"result": err,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": products,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Get All Data
func (idb *InDB) GetProducts(c *gin.Context) {
	search := c.Query("search")
	searchName := "%" + search + "%"
	var (
		products []structs.Products
		result   gin.H
	)

	idb.DB.Where("name LIKE ?", searchName).Find(&products)
	if len(products) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": products,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Create user
func (idb *InDB) CreateProduct(c *gin.Context) {
	var (
		products structs.Products
		result   gin.H
	)

	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idb.DB.Create(&products)
	result = gin.H{
		"result": products,
	}
	c.JSON(http.StatusOK, result)
}

// Update user
func (idb *InDB) UpdateProduct(c *gin.Context) {
	id := c.Query("id")

	var (
		products    structs.Products
		newProducts structs.Products
		result      gin.H
	)

	err := idb.DB.First(&products, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}

	if err := c.ShouldBindJSON(&newProducts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = idb.DB.Model(&products).Updates(newProducts).Error

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
func (idb *InDB) DeleteProduct(c *gin.Context) {
	var (
		products structs.Products
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&products, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}
	err = idb.DB.Delete(&products).Error
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
