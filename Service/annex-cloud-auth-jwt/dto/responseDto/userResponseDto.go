package reponsedto

import (

	"github.com/google/uuid"
)

type UsersResponseDto struct {
	UserId       uuid.UUID    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleName string `json:"role_name"`
	CreatedBy string `json:"created_by"`
}

type LoginResponseDto struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}