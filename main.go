package main

import "github.com/gin-gonic/gin"

func health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "okk"})
}

func main() {
	r := gin.Default()
	//hello world
	r.GET("/health", health)
	r.Run(":8080")
}
