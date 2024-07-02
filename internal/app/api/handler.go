package api

import (
	"github.com/gocql/gocql"
	"github.com/srivalli123125/Todo-Space/internal/app/models"
	"github.com/srivalli123125/Todo-Space/internal/app/repository"
)

type TodoAPI struct {
	repo repository.TodoRepository
}

func NewTodoAPI(repo repository.TodoRepository) *TodoAPI {
	return &TodoAPI{
		repo: repo,
	}
}

func (api *TodoAPI) AddTodo(todo models.Todo) error {
	return api.repo.Create(todo)
}

func (api *TodoAPI) UpdateTodo(todo models.Todo) error {
	return api.repo.Update(todo)
}

func (api *TodoAPI) DeleteTodo(id gocql.UUID) error {
	return api.repo.Delete(id)
}

func (api *TodoAPI) GetTodo(id gocql.UUID) (models.Todo, error) {
	return api.repo.GetByID(id)
}

func (api *TodoAPI) GetTodoList(page int, pageSize int, lastPageToken string, status string) ([]models.Todo, error) {
	return api.repo.GetAllList(page, pageSize, lastPageToken, status)
}
