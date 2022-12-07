package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"appd.demo/models"
)

func main() {
	req := gin.Default()

	req.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	models.ConnectDatabase()

	req.Run()
}