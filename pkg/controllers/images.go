// controllers/images.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imageservice/pkg/models"
)

type CreateImageInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"discription"`
}

type UpdateImageInput struct {
	Title       string `json:"title"`
	Description string `json:"discription"`
}

// GET /images
// Get all images
func FindImages(c *gin.Context) {
	var images []models.Image
	models.DB.Find(&images)

	c.JSON(http.StatusOK, gin.H{"data": images})
}

// GET /images/:id
// Find a image
func GetImage(c *gin.Context) { // Get model if exist
	var image models.Image

	if err := models.DB.Where("id = ?", c.Param("id")).First(&image).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": image})
}

// POST /images
// Create new image
func CreateImage(c *gin.Context) {
	// Validate input
	var input CreateImageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create image
	image := models.Image{Title: input.Title, Description: input.Description}
	models.DB.Create(&image)

	c.JSON(http.StatusOK, gin.H{"data": image})
}

// PATCH /images/:id
// Update a image
func UpdateImage(c *gin.Context) {
	// Get model if exist
	var image models.Image
	if err := models.DB.Where("id = ?", c.Param("id")).First(&image).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateImageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&image).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": image})
}
