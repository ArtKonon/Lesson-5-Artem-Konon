package handler

import "github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/service"

// CalculatorHandler exposes calculator operations to the application layer.
type CalculatorHandler struct {
	service *service.CalculatorService
}

// NewCalculatorHandler creates a calculator handler.
func NewCalculatorHandler(svc *service.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{service: svc}
}

// Add handles the add operation.
func (h *CalculatorHandler) Add(a, b float64) float64 {
	return h.service.Add(a, b)
}

// Subtract handles the subtract operation.
func (h *CalculatorHandler) Subtract(a, b float64) float64 {
	return h.service.Subtract(a, b)
}

// Multiply handles the multiply operation.
func (h *CalculatorHandler) Multiply(a, b float64) float64 {
	return h.service.Multiply(a, b)
}

// Divide handles the divide operation.
func (h *CalculatorHandler) Divide(a, b float64) (float64, error) {
	return h.service.Divide(a, b)
}
