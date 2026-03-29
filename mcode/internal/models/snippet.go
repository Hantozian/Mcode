package models

// Snippet представляет код-сниппет
type Snippet struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	LanguageID uint     `gorm:"not null;index" json:"language_id"`
	Title      string   `gorm:"type:text;not null" json:"title"`
	Filename   string   `gorm:"type:text;not null" json:"filename"`
	Content    string   `gorm:"type:text;not null" json:"content"`
	CreatedAt  int64    `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt  int64    `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Language   Language `gorm:"foreignKey:LanguageID" json:"language,omitempty"`
}

// ValidateSnippet проверяет корректность данных Snippet
func (s *Snippet) ValidateSnippet() error {
	if s.Title == "" {
		return ErrInvalidSnippetTitle
	}
	if s.Filename == "" {
		return ErrInvalidFilename
	}
	if s.Content == "" {
		return ErrInvalidContent
	}
	if s.LanguageID == 0 {
		return ErrInvalidLanguageID
	}
	return nil
}

// SnippetResponse используется для ответов API без хранения полного контента в списках
type SnippetResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Filename  string `json:"filename"`
	CreatedAt int64  `json:"created_at"`
}

// SnippetDetailResponse используется для ответов API с полным контентом
type SnippetDetailResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Filename  string `json:"filename"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// ConvertToSnippetResponse преобразует Snippet в SnippetResponse
func (s *Snippet) ConvertToSnippetResponse() SnippetResponse {
	return SnippetResponse{
		ID:        s.ID,
		Title:     s.Title,
		Filename:  s.Filename,
		CreatedAt: s.CreatedAt,
	}
}

// ConvertToSnippetDetailResponse преобразует Snippet в SnippetDetailResponse
func (s *Snippet) ConvertToSnippetDetailResponse() SnippetDetailResponse {
	return SnippetDetailResponse{
		ID:        s.ID,
		Title:     s.Title,
		Filename:  s.Filename,
		Content:   s.Content,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}
