package models

// Language представляет язык программирования
type Language struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	Slug      string    `gorm:"type:text;unique;not null" json:"slug"`
	CreatedAt int64     `gorm:"autoCreateTime:milli" json:"created_at"`
	Snippets  []Snippet `gorm:"foreignKey:LanguageID" json:"snippets,omitempty"`
}

// ValidateLanguage проверяет корректность данных Language
func (l *Language) ValidateLanguage() error {
	if l.Name == "" {
		return ErrInvalidLanguageName
	}
	if l.Slug == "" {
		return ErrInvalidLanguageSlug
	}
	return nil
}
