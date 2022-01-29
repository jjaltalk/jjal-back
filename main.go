package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	router.GET("/hongseok", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello": "Hongseok",
		})
	})

	router.Run(":5000")
}
