package repository

import (
	"errors"
	"fmt"

	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/models"
)

// ErrDivisionByZero is returned when the divisor is zero.
var ErrDivisionByZero = errors.New("calculator: division by zero")

// CalculatorRepository handles arithmetic data access and domain operations.
type CalculatorRepository struct{}

// NewCalculatorRepository creates a repository instance.
func NewCalculatorRepository() *CalculatorRepository {
	return &CalculatorRepository{}
}

// Add returns the sum of the two values.
func (r *CalculatorRepository) Add(input models.ArithmeticInput) float64 {
	return input.Left + input.Right
}

// Subtract returns the difference of the two values.
func (r *CalculatorRepository) Subtract(input models.ArithmeticInput) float64 {
	return input.Left - input.Right
}

// Multiply returns the product of the two values.
func (r *CalculatorRepository) Multiply(input models.ArithmeticInput) float64 {
	return input.Left * input.Right
}

// Divide returns the quotient of the two values.
func (r *CalculatorRepository) Divide(input models.ArithmeticInput) (float64, error) {
	if input.Right == 0 {
		return 0, fmt.Errorf("calculator: divide %v by %v: %w", input.Left, input.Right, ErrDivisionByZero)
	}
	return input.Left / input.Right, nil
}
