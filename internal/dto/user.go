package dto

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"` //If React sends {} Gin will return 400 Bad Request. This is because the binding:"required" tag tells Gin to validate that the Name field is present in the incoming JSON payload. If the field is missing or empty, Gin will consider the request invalid and respond with a 400 status code.
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
