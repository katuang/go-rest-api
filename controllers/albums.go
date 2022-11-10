package controllers

import (
	"example/go-rest-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateAlbumInput struct {
	Title	string	`json:"title" binding:"required"`
	Artist	string	`json:"artist" binding:"required"`
}

func GetAlbums(c *gin.Context) {

	var albums []models.Album
	models.DB.Find(&albums)
	c.JSON(http.StatusOK, gin.H{"data": albums})
}

func CreateAlbum(c *gin.Context) {
	// Validate input
	var input CreateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	album := models.Album{Title: input.Title, Artist: input.Artist}
	models.DB.Create(&album)

	c.JSON(http.StatusOK, gin.H{"data": album})
}