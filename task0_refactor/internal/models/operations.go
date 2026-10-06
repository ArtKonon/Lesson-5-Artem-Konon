package models

// ArithmeticInput models the input for arithmetic operations.
type ArithmeticInput struct {
	Left  float64
	Right float64
}

// TextInput models the input for text processing.
type TextInput struct {
	Content string
}

// TextStats represents the result of a text analysis.
type TextStats struct {
	Words      int
	Characters int
}
