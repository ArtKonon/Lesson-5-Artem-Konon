package textanalyzer

import (
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/repository"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/service"
)

// ErrEmptyText is kept for backward compatibility with existing tests and callers.
var ErrEmptyText = repository.ErrEmptyText

// WordCount returns the number of words in the provided text.
func WordCount(text string) (int, error) {
	return service.NewTextService(repository.NewTextRepository()).WordCount(text)
}

// CharCount returns the number of runes in the provided text.
func CharCount(text string) int {
	return service.NewTextService(repository.NewTextRepository()).CharCount(text)
}
