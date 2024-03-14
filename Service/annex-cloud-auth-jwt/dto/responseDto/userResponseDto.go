package reponsedto

import "github.com/google/uuid"

type UsersResponseDto struct {
	Id       uuid.UUID    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role string `json:"role"`
}

type LoginResponseDto struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}