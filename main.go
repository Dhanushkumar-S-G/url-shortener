package main

import "github.com/gin-gonic/gin"


func main() {
	router := gin.Default()

	router.POST("/shorten/")

	router.Run("localhost:8000")
}