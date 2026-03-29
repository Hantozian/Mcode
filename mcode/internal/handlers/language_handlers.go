package handlers

import (
	"log"
	"mcode/snippets/internal/db"
	"mcode/snippets/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLanguages получает список всех языков программирования
// GET /languages
func GetLanguages(c *gin.Context) {
	var languages []models.Language

	result := db.DB.Find(&languages)
	if result.Error != nil {
		log.Printf("Failed to fetch languages: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.ErrDatabaseError.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, languages)
}
