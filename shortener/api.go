package shortener

import "github.com/gin-gonic/gin"

func generate(c *gin.Context) {
	shortenURL := base62()
	originalURL := c.PostForm("original_url")

	
}