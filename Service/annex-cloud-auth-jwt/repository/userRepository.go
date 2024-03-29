package repository

import (
	"errors"
	"annex-cloud-auth-jwt/entity"
	"annex-cloud-auth-jwt/dto/requestDto"
	"annex-cloud-auth-jwt/helper"
    "github.com/google/uuid"
	responsedto "annex-cloud-auth-jwt/dto/responseDto"
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
	var user []entity.User
	results := u.Db.Find(&user)
	helper.ErrorPanic(results.Error)
	return user
}

// GetById implements UsersRepository
func (u *UserRepository) GetById(userId uuid.UUID) (entity.User, error) {
	var user entity.User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Save implements UsersRepository
func (u *UserRepository) Save(user requestdto.CreateUserRequestDto) {
	result := u.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository
func (u *UserRepository) Update(user entity.User) {
	updateUser := requestdto.UpdateUserRequestDto{
		Id:       uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		RoleId: 	user.RoleId,
	}
	result := u.Db.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}

// GetByUsername implements UsersRepository
func (u *UserRepository) GetByEmail(email string) (responsedto.UsersResponseDto, error) {
	var user responsedto.UsersResponseDto

    // Call stored procedure using GORM
    result := u.Db.Raw("CALL sp_GetUserByEmail(?)", email).Scan(&user)
    if result.Error != nil {
        return user, errors.New("user not found")
    }

    return user, nil
}