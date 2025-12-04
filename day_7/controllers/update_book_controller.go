package controllers

import (
	"golang-training/day_7/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBookController(c *gin.Context, store *models.LibraryStore) {
	var (
		updatedBook models.Books
		ok          bool
	)

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook, ok = store.UpdateBook(updatedBook)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
		"book":    updatedBook,
	})
}
