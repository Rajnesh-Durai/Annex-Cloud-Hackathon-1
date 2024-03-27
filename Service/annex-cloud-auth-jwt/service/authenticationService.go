package service

import (
	"annex-cloud-auth-jwt/config"
	requestdto "annex-cloud-auth-jwt/dto/requestDto"
	"annex-cloud-auth-jwt/entity"
	"annex-cloud-auth-jwt/helper"
	"annex-cloud-auth-jwt/repository"
	"annex-cloud-auth-jwt/utils"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.IUserRepository
	Validate        *validator.Validate
}

func NewAuthenticationService(usersRepository repository.IUserRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(user requestdto.LoginRequestDto) (string, error) {
	// Find username in database
	newUser, err := a.UsersRepository.GetByEmail(user.Email)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Verify password
	if err := utils.VerifyPassword(newUser.Password, user.Password); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate CustomClaims with user details
	claims := utils.CustomClaims{
		Sub:      newUser.UserId.String(),
		Role:     newUser.RoleName,
		Username: newUser.Username,
	}

	// Generate Token
	config, _ := config.LoadConfig(".")
	token, err := utils.GenerateToken(config.TokenExpiresIn, claims, config.TokenSecret)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}


// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(user requestdto.CreateUserRequestDto) {

	hashedPassword, err := utils.HashPassword(user.Password)
	helper.ErrorPanic(err)

	newUser := entity.User{
		Id: uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
		RoleId: user.RoleId,
		CreatedOn: time.Now(),
		CreatedBy: user.Username,
	}
	a.UsersRepository.Save(newUser)
}