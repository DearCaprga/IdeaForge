package main

import "github.com/gin-gonic/gin"

func health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func main() {
	r := gin.Default()
	//aaaaa
	r.GET("/health", health)
	r.Run(":8080")
}
