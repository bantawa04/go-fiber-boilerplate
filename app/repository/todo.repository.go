package repository

import (
	"github.com/bantawao4/gofiber-boilerplate/app/model"
	"github.com/bantawao4/gofiber-boilerplate/config"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetTodos(page, perPage int, searchQuery string) ([]model.TodoModel, int64, error)
	CreateTodo(todo *model.TodoModel) (*model.TodoModel, error)
	GetTodoById(todoId string) (*model.TodoModel, error)
	UpdateTodo(todo *model.TodoModel) (*model.TodoModel, error)
	DeleteTodo(id string) error
	WithTrx(tx *gorm.DB) TodoRepository
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{
		db: config.DB.Db,
	}
}

func (r *todoRepository) WithTrx(tx *gorm.DB) TodoRepository {
	if tx == nil {
		return r
	}
	// Create a new instance instead of modifying the existing one
	newRepo := &todoRepository{
		db: tx,
	}
	return newRepo
}

func (r *todoRepository) GetTodos(page, perPage int, searchQuery string) ([]model.TodoModel, int64, error) {
	var todos []model.TodoModel
	var total int64

	query := r.db.Model(&model.TodoModel{})

	if searchQuery != "" {
		query = query.Where("title ILIKE ?", "%"+searchQuery+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated data
	offset := (page - 1) * perPage
	if err := query.Offset(offset).Limit(perPage).Scan(&todos).Error; err != nil {
		return nil, 0, err
	}

	return todos, total, nil
}

func (r *todoRepository) CreateTodo(todo *model.TodoModel) (*model.TodoModel, error) {
	err := r.db.Create(todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepository) GetTodoById(todoId string) (*model.TodoModel, error) {
	var todo model.TodoModel
	err := r.db.Model(&model.TodoModel{}).Where("id = ?", todoId).First(&todo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) UpdateTodo(todo *model.TodoModel) (*model.TodoModel, error) {
	err := r.db.Model(&model.TodoModel{}).Where("id = ?", todo.ID).Updates(todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepository) DeleteTodo(id string) error {
	return r.db.Delete(&model.TodoModel{}, "id = ?", id).Error
}
