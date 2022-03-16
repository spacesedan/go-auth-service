package dto

type FindOne struct {
	Email *string `json:"email" binding:"required, email"`
}
