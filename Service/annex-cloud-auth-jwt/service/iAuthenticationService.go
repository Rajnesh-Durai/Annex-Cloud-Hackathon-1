package service

import requestdto "annex-cloud-auth-jwt/dto/requestDto"

type AuthenticationService interface {
	Login(user requestdto.LoginRequestDto) (string, error)
	Register(user requestdto.CreateUserRequestDto)
}