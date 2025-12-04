package controllers

import (
	"golang-training/day_7/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBookController(c *gin.Context, store *models.LibraryStore) {
	var bookToDelete struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&bookToDelete); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok := store.DeleteBook(bookToDelete.ID)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
