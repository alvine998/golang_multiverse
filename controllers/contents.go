package controllers

import (
	"api/webservice-multiverse/structs"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// to get data with id
func (idb *InDB) GetContent(c *gin.Context) {

	var (
		contents structs.Contents
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id=?", id).First(&contents).Error
	if err != nil {
		result = gin.H{
			"result": err,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": contents,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Get All Data
func (idb *InDB) GetContents(c *gin.Context) {
	search := c.Query("search")
	searchName := "%" + search + "%"
	var (
		contents []structs.Contents
		result   gin.H
	)

	idb.DB.Where("title LIKE ?", searchName).Find(&contents)
	if len(contents) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": contents,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// Create data
func (idb *InDB) CreateContent(c *gin.Context) {
	var (
		contents structs.Contents
		result   gin.H
	)

	title := c.PostForm("title")
	file, err := c.FormFile("image")
	notes := c.PostForm("notes")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)

	paths := strings.Replace(file.Filename, " ", "_", -1)

	filePath := "./uploads/" + paths

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	files := "https://backend.ptmultiverse.com/uploads/" + paths

	contents.Title = title
	contents.Image = files
	contents.Notes = notes

	idb.DB.Create(&contents)
	result = gin.H{
		"result": contents,
	}
	c.JSON(http.StatusOK, result)
}

// Update user
func (idb *InDB) UpdateContent(c *gin.Context) {
	id := c.Query("id")

	var (
		contents    structs.Contents
		newContents structs.Contents
		result      gin.H
	)

	err := idb.DB.First(&contents, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}

	title := c.PostForm("title")
	file, err := c.FormFile("image")
	notes := c.PostForm("notes")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)

	paths := strings.Replace(file.Filename, " ", "_", -1)

	filePath := "./uploads/" + paths

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	files := "https://backend.ptmultiverse.com/uploads/" + paths

	newContents.Title = title
	newContents.Image = files
	newContents.Notes = notes

	err = idb.DB.Model(&contents).Updates(newContents).Error

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
func (idb *InDB) DeleteContent(c *gin.Context) {
	var (
		contents structs.Contents
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&contents, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"status": "Error",
			"Code":   http.StatusBadRequest,
		}
		c.JSON(http.StatusBadRequest, result)
	}
	err = idb.DB.Delete(&contents).Error
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
