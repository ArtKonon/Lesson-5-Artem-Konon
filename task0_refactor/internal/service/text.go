package service

import (
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/models"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/repository"
)

// TextService applies text analysis business logic.
type TextService struct {
	repository *repository.TextRepository
}

// NewTextService creates a text service.
func NewTextService(repo *repository.TextRepository) *TextService {
	return &TextService{repository: repo}
}

// WordCount returns the number of words in the input text.
func (s *TextService) WordCount(text string) (int, error) {
	return s.repository.WordCount(models.TextInput{Content: text})
}

// CharCount returns the count of runes in the trimmed input text.
func (s *TextService) CharCount(text string) int {
	return s.repository.CharCount(models.TextInput{Content: text})
}
