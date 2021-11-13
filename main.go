package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleList(c *gin.Context) {
	var loadedEmails, _ = GetAll()
	var result string

	if c.Query("password") != "abcd9090" {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Password incorrect!"))
		return
	}
	for i := range loadedEmails {
		result += loadedEmails[i].Email + "<br/>"
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(result))
}

func handleSubscribe(c *gin.Context) {
	var email Email
	log.Printf("param %s", c.Param("email"))
	email.Email = c.PostForm("email")

	id, err := Create(&email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err, "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id, "success": true})
}

func main() {
	r := gin.Default()
	r.GET("/list/", handleList)
	r.POST("/subscribe/", handleSubscribe)
	r.Run() // listen and serve on 0.0.0.0:8080
}
