package controllers

import (
	"golang-training/day_7/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBookController(c *gin.Context, store *models.LibraryStore) {
	var newBook models.Books

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newBook.Title == "" || newBook.Author == "" || newBook.Genre == "" || newBook.Rating == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title, Author, Genre, and Rating are required fields"})
		return
	}

	store.AddBook(&newBook)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Book added successfully",
		"book":    newBook,
	})
}
