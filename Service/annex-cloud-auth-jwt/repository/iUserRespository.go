package repository

import (
	"annex-cloud-auth-jwt/entity"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Save(user entity.User)
	Update(user entity.User)
	Delete(userId uuid.UUID)
	GetById(userId uuid.UUID) (entity.User, error)
	GetAll() []entity.User
	GetByEmail(email string) (entity.User, error)
}