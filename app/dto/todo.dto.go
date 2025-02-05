package dto

import (
	"time"

	"github.com/bantawao4/gofiber-boilerplate/app/model"
)

type TodoResponse struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func ToTodoResponse(todo *model.TodoModel) *TodoResponse {
	return &TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Status:    todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func ToTodoListResponse(todos []model.TodoModel) []TodoResponse {
	response := make([]TodoResponse, 0)
	for _, todo := range todos {
		response = append(response, *ToTodoResponse(&todo))
	}
	return response
}
