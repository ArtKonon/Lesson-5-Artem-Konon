package main

import (
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/app"
)

func main() {
	application := app.New()

	_ = application.Calculator.Add(2, 3)
	_, _ = application.Calculator.Divide(10, 0)
	_, _ = application.Text.WordCount("the quick brown fox")
}
