package requestdto

import (
	"github.com/google/uuid"
)

type CreateUserRequestDto struct {
	Username string `validate:"required,min=2,max=50" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
	Role string `validate:"required,min=2,max=50" json:"role"`
}

type UpdateUserRequestDto struct {
	Id       uuid.UUID    `validate:"required"`
	Username string `validate:"required,max=50,min=2" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
	Role string `validate:"required,min=2,max=50" json:"role"`
}

type LoginRequestDto struct {
	Email string `validate:"required,max=100,min=2" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}
