package service

import (
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/models"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/repository"
)

// CalculatorService applies arithmetic business logic.
type CalculatorService struct {
	repository *repository.CalculatorRepository
}

// NewCalculatorService creates a calculator service.
func NewCalculatorService(repo *repository.CalculatorRepository) *CalculatorService {
	return &CalculatorService{repository: repo}
}

// Add returns the sum of two numbers.
func (s *CalculatorService) Add(a, b float64) float64 {
	return s.repository.Add(models.ArithmeticInput{Left: a, Right: b})
}

// Subtract returns the difference between two numbers.
func (s *CalculatorService) Subtract(a, b float64) float64 {
	return s.repository.Subtract(models.ArithmeticInput{Left: a, Right: b})
}

// Multiply returns the product of two numbers.
func (s *CalculatorService) Multiply(a, b float64) float64 {
	return s.repository.Multiply(models.ArithmeticInput{Left: a, Right: b})
}

// Divide returns the quotient of two numbers.
func (s *CalculatorService) Divide(a, b float64) (float64, error) {
	return s.repository.Divide(models.ArithmeticInput{Left: a, Right: b})
}
