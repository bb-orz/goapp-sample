package dtos

type InDTO struct {
	Email string `validate:"required,email"`
}

