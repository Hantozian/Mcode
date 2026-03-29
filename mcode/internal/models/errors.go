package models

import "fmt"

// Ошибки валидации
var (
	ErrInvalidLanguageName = fmt.Errorf("language name cannot be empty")
	ErrInvalidLanguageSlug = fmt.Errorf("language slug cannot be empty")
	ErrInvalidSnippetTitle = fmt.Errorf("snippet title cannot be empty")
	ErrInvalidFilename     = fmt.Errorf("filename cannot be empty")
	ErrInvalidContent      = fmt.Errorf("content cannot be empty")
	ErrInvalidLanguageID   = fmt.Errorf("language id is required")
	ErrLanguageNotFound    = fmt.Errorf("language not found")
	ErrSnippetNotFound     = fmt.Errorf("snippet not found")
	ErrDatabaseError       = fmt.Errorf("database error")
)

// ErrorResponse стандартный формат ошибки API
type ErrorResponse struct {
	Error string `json:"error"`
}
