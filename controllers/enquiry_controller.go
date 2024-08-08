package controllers

import (
	"net/http"

	"apsdatabase/db"
	"apsdatabase/models"
	"github.com/gin-gonic/gin"
)

func CreateEnquiry(c *gin.Context) {
	var enquiry models.Enquiry
	if err := c.ShouldBindJSON(&enquiry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&enquiry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enquiry)
}

func GetEnquiry(c *gin.Context) {
	id := c.Param("id")
	var enquiry models.Enquiry
	if err := db.DB.First(&enquiry, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "enquiry not found"})
		return
	}

	c.JSON(http.StatusOK, enquiry)
}

func GetAllEnquiry(c *gin.Context) {
	var enquirys []models.Enquiry
	if err := db.DB.Find(&enquirys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enquirys)
}

func UpdateEnquiry(c *gin.Context) {
	id := c.Param("id")
	var enquiry models.Enquiry
	if err := db.DB.First(&enquiry, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "enquiry not found"})
		return
	}

	if err := c.ShouldBindJSON(&enquiry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&enquiry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enquiry)
}

func DeleteEnquiry(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Enquiry{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "enquiry deleted successfully"})
}
