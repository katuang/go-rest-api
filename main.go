package main

import (

	"github.com/gin-gonic/gin"

	"example/go-rest-api/controllers"
	"example/go-rest-api/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/endpoint", controllers.GetAlbums)
	router.POST("/albums", controllers.CreateAlbum)
	router.Run("localhost:7070")
}

