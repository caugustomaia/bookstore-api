package controllers

import (
	"bookstore-api/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	var book models.Book
	//var response string

	cache_val, err := models.RDB.Get(c.Param("id")).Result()
	if err != nil {
		if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Record not found!",
			})
			return
		}

		byte_value, err := json.Marshal(book)
		if err != nil {
			panic(err)
		}

		if err := models.RDB.Set(c.Param("id"), byte_value, 5000000000).Err(); err != nil { // 10000000000 = 10 seconds
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": book,
		})
		return
	}

	if err := json.Unmarshal([]byte(cache_val), &book); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create book
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// get model if exists
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	// validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// converter o tipo UpdateBookInput para models.Book
	b := func(a UpdateBookInput) models.Book {
		//var b models.Book
		return models.Book{
			Title:  a.Title,
			Author: a.Title,
		}

	}

	models.DB.Model(&book).Updates(b)

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

// DELETE /books/:id
// Deleta um livro
func DeleteBook(c *gin.Context) {
	// Get model if exists
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}
	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})

}
