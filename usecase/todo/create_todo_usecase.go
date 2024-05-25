package todo

import (
	"context"

	"github.com/nftug/wails-todo-app/domain/todo"
	"github.com/nftug/wails-todo-app/interfaces"
)

type CreateTodoUseCase struct {
	repo todo.TodoRepository
}

func NewCreateTodoUseCase(repo todo.TodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{repo}
}

func (u *CreateTodoUseCase) Execute(command todo.CreateCommand, ctx context.Context) (*interfaces.CreatedResponse, error) {
	t, err := todo.NewTodo(command)
	if err != nil {
		return nil, err
	}
	if err := u.repo.Save(t, ctx); err != nil {
		return nil, err
	}
	return interfaces.NewCreatedResponse(t), nil
}