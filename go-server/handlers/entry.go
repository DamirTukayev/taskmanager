package handlers

import (
	"net/http"
	"tasksmanager/database"
	"tasksmanager/models"

	"github.com/gin-gonic/gin"
)

func CreateEntry(c *gin.Context) {
	var entry models.Entry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Если необходимо, можно здесь дополнительно проверить формат даты и времени (например, через регулярные выражения)
	database.DB.Create(&entry)
	c.JSON(http.StatusCreated, entry)
}

func GetEntries(c *gin.Context) {
	var entries []models.Entry
	if err := database.DB.Find(&entries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entries)
}

func GetEntryByID(c *gin.Context) {
	var entry models.Entry
	if err := database.DB.Where("id = ?", c.Param("id")).First(&entry).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "entry not found"})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func UpdateEntry(c *gin.Context) {
	var entry models.Entry
	if err := database.DB.Where("id = ?", c.Param("id")).First(&entry).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "entry not found"})
		return
	}

	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&entry)
	c.JSON(http.StatusOK, entry)
}

func DeleteEntry(c *gin.Context) {
	if err := database.DB.Where("id = ?", c.Param("id")).Delete(&models.Entry{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "entry not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
