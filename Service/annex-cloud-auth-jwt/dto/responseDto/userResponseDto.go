package reponsedto

import (
	"time"

	"github.com/google/uuid"
)

type UsersResponseDto struct {
	Id       uuid.UUID    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role_name"`
	CreatedOn time.Time `json:"created_on"`
	CreatedBy string `json:"created_by"`
}

type LoginResponseDto struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}