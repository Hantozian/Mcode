package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"mcode/snippets/internal/db"
	"mcode/snippets/internal/handlers"
)

func main() {
	// Загружаем переменные окружения из .env файла
	if err := godotenv.Load("config/.env"); err != nil {
		log.Println("Note: .env file not found, using system environment variables")
	}

	// Инициализируем подключение к БД
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
	}
	log.Println("✅ Database initialized successfully")

	// Отложить закрытие подключения
	defer func() {
		if err := db.CloseDB(); err != nil {
			log.Printf("Error closing database: %v\n", err)
		}
	}()

	// Создаем маршрутизатор Gin
	r := gin.Default()

	// Добавляем CORS middleware для будущего фронтенда
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// ============ МАРШРУТЫ ============

	// Языки программирования
	r.GET("/languages", handlers.GetLanguages)

	// Сниппеты
	r.POST("/languages/:slug/snippets", handlers.CreateSnippet)
	r.GET("/languages/:slug/snippets", handlers.GetSnippets)
	r.GET("/languages/:slug/snippets/:id/content", handlers.GetSnippetContent)
	r.DELETE("/languages/:slug/snippets/:id", handlers.DeleteSnippet)

	// Здоровье приложения (для проверки, что сервер запущен)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Запускаем сервер
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("\n🚀 Server starting on http://localhost:%s\n", port)
	log.Printf("📚 API available at http://localhost:%s/\n\n", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
