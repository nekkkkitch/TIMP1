package scripts

import (
	"context"
	"window/pkg/models"
)

type Tabler interface {
	InitTable() []models.Lesson
	AddElement(string) (models.Lesson, error)
	SaveTable([][]string) []error
}

type App struct {
	ctx    context.Context
	tabler Tabler
}

func NewApp(tabler Tabler) *App {
	return &App{tabler: tabler}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Startup(ctx context.Context) {
	a.startup(ctx)
}

func (a *App) InitTable() []models.Lesson {
	return a.tabler.InitTable()
}

func (a *App) AddElement(input string) (models.Lesson, error) {
	return a.tabler.AddElement(input)
}

func (a *App) SaveTable(table [][]string) []error {
	return a.tabler.SaveTable(table)
}
