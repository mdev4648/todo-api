package dto

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required,max=200"`
	Description string `json:"description"`
}
