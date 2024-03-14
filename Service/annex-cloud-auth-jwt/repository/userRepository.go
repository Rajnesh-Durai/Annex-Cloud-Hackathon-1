package repository

import (
	"errors"
	"annex-cloud-auth-jwt/entity"
	"annex-cloud-auth-jwt/dto/requestDto"
	"annex-cloud-auth-jwt/helper"
    "github.com/google/uuid"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) IUserRepository {
	return &UserRepository{Db: Db}
}

// Delete implements UsersRepository
func (u *UserRepository) Delete(userId uuid.UUID) {
	var user entity.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// GetAll implements UsersRepository
func (u *UserRepository) GetAll() []entity.User {
	var users []entity.User
	results := u.Db.Find(&users)
	helper.ErrorPanic(results.Error)
	return users
}

// GetById implements UsersRepository
func (u *UserRepository) GetById(usersId uuid.UUID) (entity.User, error) {
	var user entity.User
	result := u.Db.Find(&user, usersId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("users is not found")
	}
}

// Save implements UsersRepository
func (u *UserRepository) Save(user entity.User) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository
func (u *UserRepository) Update(user entity.User) {
	updateUser := requestdto.UpdateUsersRequestDto{
		Id:       uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role: 	user.Role,
	}
	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

// GetByUsername implements UsersRepository
func (u *UserRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	result := u.Db.First(&user, "email = ?", email)

	if result.Error != nil {
		return user, errors.New("invalid email or Password")
	}
	return user, nil
}