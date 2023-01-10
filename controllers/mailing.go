package controllers

import (
	"api/webservice-multiverse/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
)

// Send Email
func (idb *InDB) SendEmail(c *gin.Context) {
	var (
		mailing structs.Mailing
		result  gin.H
	)

	if err := c.ShouldBindJSON(&mailing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "vira@ptmultiverse.com")
	m.SetHeader("To", mailing.Email)
	m.SetHeader("Subject", mailing.Subject)
	m.SetBody("text/html", mailing.Message)

	d := gomail.NewDialer("mail.ptmultiverse.com", 587, "vira@ptmultiverse.com", "Multiverse888")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}

	idb.DB.Create(&mailing)

	result = gin.H{
		"result":   "Success Sending Mail",
		"creating": mailing,
	}
	c.JSON(http.StatusOK, result)
}
