package app

import (
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/handler"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/repository"
	"github.com/softserve/go-with-genai-topic4-error-handling/task0_refactor/internal/service"
)

// App wires together repository, service, and handler layers.
type App struct {
	Calculator *handler.CalculatorHandler
	Text       *handler.TextHandler
}

// New creates the application with initialized dependencies.
func New() *App {
	calcRepo := repository.NewCalculatorRepository()
	textRepo := repository.NewTextRepository()

	calcSvc := service.NewCalculatorService(calcRepo)
	textSvc := service.NewTextService(textRepo)

	return &App{
		Calculator: handler.NewCalculatorHandler(calcSvc),
		Text:       handler.NewTextHandler(textSvc),
	}
}
