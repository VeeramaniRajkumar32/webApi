package main

import (
	"github.com/VeeramaniRajkumar32/webApi/include"
	"github.com/gin-gonic/gin"
)

func init() {
	include.LoadEnvVariables()
	include.Conn()
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
