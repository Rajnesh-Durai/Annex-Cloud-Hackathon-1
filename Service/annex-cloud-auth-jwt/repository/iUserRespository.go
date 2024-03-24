package repository

import (
	"annex-cloud-auth-jwt/entity"
	"github.com/google/uuid"
	"annex-cloud-auth-jwt/dto/requestDto"
	responsedto "annex-cloud-auth-jwt/dto/responseDto"
)

type IUserRepository interface {
	Save(user requestdto.CreateUserRequestDto)
	Update(user entity.User)
	Delete(userId uuid.UUID)
	GetById(userId uuid.UUID) (entity.User, error)
	GetAll() []entity.User
	GetByEmail(email string) (responsedto.UsersResponseDto, error)
}
