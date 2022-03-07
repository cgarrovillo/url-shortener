package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"url-shortener/handler"
	"url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "URL Shortener API",
		})
	})

	// Creating short urls
	r.POST("/create", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// retrieving original urls.
	// Param is tightly coupled with HandleShortUrlRedirect underneath
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}