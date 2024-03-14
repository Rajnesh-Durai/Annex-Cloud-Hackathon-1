package service

import requestdto "annex-cloud-auth-jwt/dto/requestDto"

type AuthenticationService interface {
	Login(users requestdto.LoginRequestDto) (string, error)
	Register(users requestdto.CreateUsersRequestDto)
}