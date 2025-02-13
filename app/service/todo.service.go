package service

import (
	"math"

	"github.com/bantawao4/gofiber-boilerplate/app/errors"
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/app/repository"
	"github.com/bantawao4/gofiber-boilerplate/app/response"
	"gorm.io/gorm"
)

type TodoService interface {
	GetTodos(page, perPage int, searchQuery string) ([]model.TodoModel, *response.PaginationMeta, error)
	CreateTodo(todo *model.TodoModel) (*model.TodoModel, error)
	GetTodoById(id string) (*model.TodoModel, error)
	UpdateTodo(id string, todo *model.TodoModel) (*model.TodoModel, error)
	DeleteTodo(id string) error
	WithTrx(tx *gorm.DB) TodoService
}

type todoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{todoRepo: todoRepo}
}

func (s *todoService) WithTrx(tx *gorm.DB) TodoService {
	newService := &todoService{
		todoRepo: s.todoRepo.WithTrx(tx),
	}
	return newService
}

func (s *todoService) GetTodos(page, perPage int, searchQuery string) ([]model.TodoModel, *response.PaginationMeta, error) {
	todos, total, err := s.todoRepo.GetTodos(page, perPage, searchQuery)
	if err != nil {
		return nil, nil, errors.NewInternalError(err)
	}

	if len(todos) == 0 {
		return todos, &response.PaginationMeta{
			Page:       page,
			PerPage:    perPage,
			TotalPages: 0,
			TotalItems: 0,
		}, nil
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	meta := &response.PaginationMeta{
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
		TotalItems: total,
	}

	return todos, meta, nil
}

func (s *todoService) CreateTodo(todo *model.TodoModel) (*model.TodoModel, error) {
	if todo == nil {
		return nil, errors.NewBadRequestError("Todo data cannot be empty")
	}

	createdTodo, err := s.todoRepo.CreateTodo(todo)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return createdTodo, nil
}

func (s *todoService) GetTodoById(id string) (*model.TodoModel, error) {
	if id == "" {
		return nil, errors.NewBadRequestError("Todo ID cannot be empty")
	}

	todo, err := s.todoRepo.GetTodoById(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if todo == nil {
		return nil, errors.NewNotFoundError("Todo not found")
	}
	return todo, nil
}

func (s *todoService) UpdateTodo(id string, updateData *model.TodoModel) (*model.TodoModel, error) {
	existingTodo, err := s.todoRepo.GetTodoById(id)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	if existingTodo == nil {
		return nil, errors.NewNotFoundError("Todo not found")
	}

	if updateData.Title != "" {
		existingTodo.Title = updateData.Title
	}
	if updateData.Status != "" {
		existingTodo.Status = updateData.Status
	}

	return s.todoRepo.UpdateTodo(existingTodo)
}

func (s *todoService) DeleteTodo(id string) error {
	existingTodo, err := s.todoRepo.GetTodoById(id)
	if err != nil {
		return errors.NewInternalError(err)
	}
	if existingTodo == nil {
		return errors.NewNotFoundError("Todo not found")
	}

	return s.todoRepo.DeleteTodo(id)
}
