package controllers

import (
	"golang-training/day_7/models"

	"github.com/gin-gonic/gin"
)

func GetBooksController(c *gin.Context, store *models.LibraryStore) {
	booksList := store.GetAllBooks()
	c.JSON(200, gin.H{
		"message": "Books retrieved successfully",
		"books":   booksList,
	})
}
