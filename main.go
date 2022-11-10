package main

import (

	"github.com/gin-gonic/gin"

	"example/go-rest-api/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/endpoint", controllers.GetAlbums)
	router.POST("/albums", controllers.CreateAlbum)
	router.Run("localhost:8080")
}

