package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/models"
)

// ErrEmptyText is returned when the provided text is empty.
var ErrEmptyText = errors.New("textanalyzer: empty text")

// TextRepository handles text data access and domain operations.
type TextRepository struct{}

// NewTextRepository creates a repository instance.
func NewTextRepository() *TextRepository {
	return &TextRepository{}
}

// WordCount returns the number of words in the text.
func (r *TextRepository) WordCount(input models.TextInput) (int, error) {
	trimmed := strings.TrimSpace(input.Content)
	if trimmed == "" {
		return 0, fmt.Errorf("textanalyzer: word count: %w", ErrEmptyText)
	}
	return len(strings.Fields(trimmed)), nil
}

// CharCount returns the count of runes in the trimmed text without surrounding spaces.
func (r *TextRepository) CharCount(input models.TextInput) int {
	trimmed := strings.TrimSpace(input.Content)
	if trimmed == "" {
		return 0
	}
	return len([]rune(trimmed))
}
