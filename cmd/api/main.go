package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "I'm alive",
		})
	})
	CategoryRoutes(router)
	router.Run(":8000")
}
