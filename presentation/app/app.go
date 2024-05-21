package app

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/presentation/types/dialog"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) (string, error) {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return *new(string), errors.New("name is empty")
	}

	return fmt.Sprintf("Hello %s, It's show time!", name), nil
}

func (a *App) ShowMessageDialog(opt dialog.DialogOptions) dialog.DialogButton {
	ret, _ := runtime.MessageDialog(a.ctx, opt.ToRuntimeOptions())
	return dialog.GetDialogButtonResult(ret)
}

func (a *App) CreateTodo(command todo.CreateCommand) {
	todo.NewTodo(command)
}