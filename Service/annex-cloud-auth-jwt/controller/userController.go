package controller

import (
	responsedto "annex-cloud-auth-jwt/dto/responseDto"
	"annex-cloud-auth-jwt/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.IUserRepository
}

func NewUsersController(repository repository.IUserRepository) *UserController {
	return &UserController{userRepository: repository}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	// Check if the current user is an admin
	currentRole, _ := ctx.Get("currentRole")
	if currentRole != "Admin" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to access this resource"})
		return
	}

	// Fetch users data
	users := controller.userRepository.GetAll()
	webResponse := responsedto.ResponseDto{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
