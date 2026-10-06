package handler

import "github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/service"

// TextHandler exposes text analysis operations to the application layer.
type TextHandler struct {
	service *service.TextService
}

// NewTextHandler creates a text handler.
func NewTextHandler(svc *service.TextService) *TextHandler {
	return &TextHandler{service: svc}
}

// WordCount handles a word count request.
func (h *TextHandler) WordCount(text string) (int, error) {
	return h.service.WordCount(text)
}

// CharCount handles a character count request.
func (h *TextHandler) CharCount(text string) int {
	return h.service.CharCount(text)
}
