package db

import (
	"fmt"
	"log"
	"mcode/snippets/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB инициализирует подключение к PostgreSQL и выполняет миграции
func InitDB() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v\n", err)
		return err
	}

	DB = database
	log.Println("✅ Database connection established")

	// Выполняем автомиграции
	err = DB.AutoMigrate(&models.Language{}, &models.Snippet{})
	if err != nil {
		log.Printf("Failed to run migrations: %v\n", err)
		return err
	}

	log.Println("✅ Migrations completed")

	// Заполняем начальные данные (языки программирования)
	err = seedLanguages()
	if err != nil {
		log.Printf("Warning: Failed to seed languages: %v\n", err)
	}

	return nil
}

// seedLanguages добавляет начальные языки программирования в БД
func seedLanguages() error {
	languages := []models.Language{
		{Name: "Go", Slug: "go"},
		{Name: "Python", Slug: "python"},
		{Name: "JavaScript", Slug: "javascript"},
		{Name: "Java", Slug: "java"},
		{Name: "C++", Slug: "cpp"},
		{Name: "C#", Slug: "csharp"},
		{Name: "Ruby", Slug: "ruby"},
		{Name: "PHP", Slug: "php"},
		{Name: "Rust", Slug: "rust"},
		{Name: "TypeScript", Slug: "typescript"},
	}

	for _, lang := range languages {
		// Проверяем, существует ли уже язык с таким slug
		var count int64
		DB.Model(&models.Language{}).Where("slug = ?", lang.Slug).Count(&count)
		if count == 0 {
			result := DB.Create(&lang)
			if result.Error != nil {
				return result.Error
			}
			log.Printf("✅ Added language: %s\n", lang.Name)
		}
	}

	return nil
}

// CloseDB закрывает подключение к БД
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
