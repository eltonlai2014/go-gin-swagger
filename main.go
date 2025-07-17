package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 將 static/swagger 目錄作為靜態網站提供
	r.Static("/swagger", "./static/swagger")

	// 測試用 API
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
