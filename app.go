package main

import (
	"context"
	"fmt"

	"name-for-me/ollama"
)

// App struct
type App struct {
	ctx    context.Context
	client *ollama.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	cli, _ := ollama.NewClient("http://localhost:11434")
	return &App{
		client: cli,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetModels returns a list of models
func (a *App) GetModels() ([]string, error) {
	models, err := a.client.GetModels(a.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get models: %w", err)
	}

	return models, nil
}

// GetLanguages returns a list of languages
func (a *App) GetLanguages() []string {
	return ollama.GetLanguageList()
}

// Generate generates a name for the given model and input
func (a *App) Generate(model string, language string, input string) (string, error) {
	switch language {
	case ollama.LanguageGo:
		input = ollama.MakeGoPrompt(input)
	case ollama.LanguagePython:
	case ollama.LanguageTypescript:
	}

	var result string
	err := a.client.Generate(a.ctx, model, input, func(content string, done bool) {
		result += content
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate name: %w", err)
	}

	return result, nil
}
