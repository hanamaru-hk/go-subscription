package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func handleList(c *gin.Context) {
	var loadedEmails, _ = GetAll()
	var result string

	if c.Query("password") != goDotEnvVariable("LIST_PASSWORD") {
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
	r.Use(Cors())
	r.GET("/list/", handleList)
	r.POST("/subscribe/", handleSubscribe)
	r.Run(goDotEnvVariable("PORT")) // listen and serve on 0.0.0.0:4001
}
