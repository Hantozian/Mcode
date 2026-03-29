package handlers

import (
	"log"
	"mcode/snippets/internal/db"
	"mcode/snippets/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSnippetRequest используется для распарсивания JSON-тела запроса
type CreateSnippetRequest struct {
	Title    string `json:"title" binding:"required"`
	Filename string `json:"filename" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

// CreateSnippet создает новый сниппет
// POST /languages/:slug/snippets
func CreateSnippet(c *gin.Context) {
	slug := c.Param("slug")

	// Находим язык по slug
	var language models.Language
	result := db.DB.Where("slug = ?", slug).First(&language)
	if result.Error != nil {
		log.Printf("Language not found: %s\n", slug)
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: models.ErrLanguageNotFound.Error(),
		})
		return
	}

	// Парсим JSON запрос
	var req CreateSnippetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid request body: %v\n", err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Создаем новый сниппет
	snippet := models.Snippet{
		LanguageID: language.ID,
		Title:      req.Title,
		Filename:   req.Filename,
		Content:    req.Content,
	}

	// Валидируем сниппет
	if err := snippet.ValidateSnippet(); err != nil {
		log.Printf("Snippet validation error: %v\n", err)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// Сохраняем в БД
	result = db.DB.Create(&snippet)
	if result.Error != nil {
		log.Printf("Failed to create snippet: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.ErrDatabaseError.Error(),
		})
		return
	}

	log.Printf("✅ Snippet created: %s\n", snippet.Title)
	c.JSON(http.StatusCreated, snippet.ConvertToSnippetDetailResponse())
}

// GetSnippets получает список сниппетов для конкретного языка
// GET /languages/:slug/snippets
func GetSnippets(c *gin.Context) {
	slug := c.Param("slug")

	// Находим язык по slug
	var language models.Language
	result := db.DB.Where("slug = ?", slug).First(&language)
	if result.Error != nil {
		log.Printf("Language not found: %s\n", slug)
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: models.ErrLanguageNotFound.Error(),
		})
		return
	}

	// Получаем все сниппеты для этого языка
	var snippets []models.Snippet
	result = db.DB.Where("language_id = ?", language.ID).Find(&snippets)
	if result.Error != nil {
		log.Printf("Failed to fetch snippets: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.ErrDatabaseError.Error(),
		})
		return
	}

	// Преобразуем в SnippetResponse (без полного контента)
	responses := make([]models.SnippetResponse, len(snippets))
	for i, s := range snippets {
		responses[i] = s.ConvertToSnippetResponse()
	}

	c.JSON(http.StatusOK, responses)
}

// GetSnippetContent получает полное содержимое конкретного сниппета
// GET /languages/:slug/snippets/:id/content
func GetSnippetContent(c *gin.Context) {
	slug := c.Param("slug")
	idStr := c.Param("id")

	// Парсим ID сниппета
	snippetID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("Invalid snippet ID: %s\n", idStr)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid snippet ID",
		})
		return
	}

	// Находим язык по slug
	var language models.Language
	result := db.DB.Where("slug = ?", slug).First(&language)
	if result.Error != nil {
		log.Printf("Language not found: %s\n", slug)
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: models.ErrLanguageNotFound.Error(),
		})
		return
	}

	// Получаем сниппет
	var snippet models.Snippet
	result = db.DB.Where("id = ? AND language_id = ?", snippetID, language.ID).First(&snippet)
	if result.Error != nil {
		log.Printf("Snippet not found: id=%d, language=%s\n", snippetID, slug)
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: models.ErrSnippetNotFound.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snippet.ConvertToSnippetDetailResponse())
}

// DeleteSnippet удаляет сниппет
// DELETE /languages/:slug/snippets/:id
func DeleteSnippet(c *gin.Context) {
	slug := c.Param("slug")
	idStr := c.Param("id")

	// Парсим ID сниппета
	snippetID, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("Invalid snippet ID: %s\n", idStr)
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid snippet ID",
		})
		return
	}

	// Находим язык по slug
	var language models.Language
	result := db.DB.Where("slug = ?", slug).First(&language)
	if result.Error != nil {
		log.Printf("Language not found: %s\n", slug)
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: models.ErrLanguageNotFound.Error(),
		})
		return
	}

	// Получаем сниппет для проверки существования
	var snippet models.Snippet
	result = db.DB.Where("id = ? AND language_id = ?", snippetID, language.ID).First(&snippet)
	if result.Error != nil {
		log.Printf("Snippet not found: id=%d, language=%s\n", snippetID, slug)
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: models.ErrSnippetNotFound.Error(),
		})
		return
	}

	// Удаляем сниппет
	result = db.DB.Delete(&snippet)
	if result.Error != nil {
		log.Printf("Failed to delete snippet: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.ErrDatabaseError.Error(),
		})
		return
	}

	log.Printf("✅ Snippet deleted: %s\n", snippet.Title)
	c.JSON(http.StatusOK, gin.H{"message": "Snippet deleted successfully"})
}
