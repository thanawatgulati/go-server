package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func brokenHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "James"})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Gin🍹",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", helloHandler)
	r.GET("/broken", brokenHandler)
	r.Run()
}
