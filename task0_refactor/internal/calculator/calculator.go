package calculator

import (
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/repository"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/service"
)

// ErrDivisionByZero is kept for backward compatibility with existing tests and callers.
var ErrDivisionByZero = repository.ErrDivisionByZero

// Add returns the sum of the provided values.
func Add(a, b float64) float64 {
	return service.NewCalculatorService(repository.NewCalculatorRepository()).Add(a, b)
}

// Subtract returns the difference of the provided values.
func Subtract(a, b float64) float64 {
	return service.NewCalculatorService(repository.NewCalculatorRepository()).Subtract(a, b)
}

// Multiply returns the product of the provided values.
func Multiply(a, b float64) float64 {
	return service.NewCalculatorService(repository.NewCalculatorRepository()).Multiply(a, b)
}

// Divide returns the quotient of the provided values.
func Divide(a, b float64) (float64, error) {
	return service.NewCalculatorService(repository.NewCalculatorRepository()).Divide(a, b)
}
