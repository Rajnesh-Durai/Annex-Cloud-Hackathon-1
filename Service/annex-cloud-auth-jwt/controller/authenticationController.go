package controller

import (
	"fmt"
	responsedto "annex-cloud-auth-jwt/dto/responseDto"
	requestdto "annex-cloud-auth-jwt/dto/requestDto"
	"annex-cloud-auth-jwt/helper"
	"annex-cloud-auth-jwt/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: service}
}

func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := requestdto.LoginRequestDto{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)

	token, err_token := controller.authenticationService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := responsedto.ResponseDto{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid email or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := responsedto.LoginResponseDto{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := responsedto.ResponseDto{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUserRequest := requestdto.CreateUserRequestDto{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	controller.authenticationService.Register(createUserRequest)

	webResponse := responsedto.ResponseDto{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
