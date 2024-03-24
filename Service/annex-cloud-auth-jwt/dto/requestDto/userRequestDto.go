package requestdto

import (
	"github.com/google/uuid"
	"time"
)

type CreateUserRequestDto struct {
	Id       uuid.UUID `validate:"required" json:"user_id"`
	Username string `validate:"required,min=2,max=50" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
	RoleId    uuid.UUID    `validate:"required" json:"role_id"`
	CreatedOn time.Time `validate:"required" json:"created_on"`
	CreatedBy string `validate:"required,min=2,max=50" json:"created_by"`
}

type UpdateUserRequestDto struct {
	Id       uuid.UUID `validate:"required" json:"user_id"`
	Username string    `validate:"required,max=50,min=2" json:"username"`
	Email    string    `validate:"required,min=2,max=100" json:"email"`
	Password string    `validate:"required,min=2,max=100" json:"password"`
	RoleId   uuid.UUID       `validate:"required" json:"role_id"`
}

type LoginRequestDto struct {
	Email    string `validate:"required,max=100,min=2" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}
