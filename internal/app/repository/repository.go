package repository

import (
	"github.com/gocql/gocql"
	"github.com/srivalli123125/Todo-Space/internal/app/repository"
)

type TodoRepository interface {
	Create(todo models.Todo) error
	Update(todo models.Todo) error
	Delete(id gocql.UUID) error
	GetByID(id gocql.UUID) (models.Todo, error)
	GetAllList(page int, pageSize int, lastPageToken string, status string) ([]models.Todo, error)
}
type scyllaTodoRepository struct {
	session *gocql.Session
}

func NewTodoRepository(session *gocql.Session) TodoRepository {
	return &scyllaTodoRepository{
		session: session,
	}
}

func (repo *scyllaTodoRepository) Create(todo models.Todo) error {
	query := "INSERT INTO todos (id, user_id, title, description, status, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?)"
	err := repo.session.Query(query,
		todo.ID,
		todo.UserID,
		todo.Title,
		todo.Description,
		todo.Status,
		todo.Created,
		todo.Updated,
	).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (repo *scyllaTodoRepository) Update(todo models.Todo) error {
	query := "UPDATE todos SET title = ?, description = ?, status = ?, updated = ? WHERE id = ? AND user_id = ? "
	err := repo.session.Query(query,
		todo.Title,
		todo.Description,
		todo.Status,
		todo.Updated,
		todo.ID,
		todo.UserID,
	).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (repo *scyllaTodoRepository) Delete(id gocql.UUID) error {
	query := "DELETE FROM todos WHERE id = ?"
	err := repo.session.Query(query, id).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (repo *scyllaTodoRepository) GetByID(id gocql.UUID) (models.Todo, error) {
	var todo models.Todo
	query := "SELECT id, user_id, title, description, status, created, updated FROM todos WHERE id = ? LIMIT 1"
	if err := repo.session.Query(query, id).Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Description,
		&todo.Status,
		&todo.Created,
		&todo.Updated,
	); err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (repo *scyllaTodoRepository) GetAllList(page int, pageSize int, lastPageToken string, status string) ([]models.Todo, error) {
	var todos []models.Todo
	var query string
	var iter *gocql.Iter

	if status != "" {
		query = "SELECT id, user_id, title, description, status, created, updated FROM todos WHERE status = ? LIMIT ? ALLOW FILTERING"
		if lastPageToken != "" {
			query = "SELECT id, user_id, title, description, status, created, updated FROM todos WHERE status = ? AND token(id) > token(?) LIMIT ? ALLOW FILTERING"
			iter = repo.session.Query(query, status, lastPageToken, pageSize).Iter()
		} else {
			iter = repo.session.Query(query, status, pageSize).Iter()
		}
	} else {
		query = "SELECT id, user_id, title, description, status, created, updated FROM todos LIMIT ? ALLOW FILTERING"
		if lastPageToken != "" {
			query = "SELECT id, user_id, title, description, status, created, updated FROM todos WHERE token(id) > token(?) LIMIT ? ALLOW FILTERING"
			iter = repo.session.Query(query, lastPageToken, pageSize).Iter()
		} else {
			iter = repo.session.Query(query, pageSize).Iter()
		}
	}

	var todo models.Todo
	for iter.Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Description,
		&todo.Status,
		&todo.Created,
		&todo.Updated,
	) {
		todos = append(todos, todo)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}
	return todos, nil
}
