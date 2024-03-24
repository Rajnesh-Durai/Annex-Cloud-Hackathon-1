package main

import (
	"annex-cloud-auth-jwt/config"
	"annex-cloud-auth-jwt/controller"
	"annex-cloud-auth-jwt/entity"
	"annex-cloud-auth-jwt/helper"
	"annex-cloud-auth-jwt/repository"
	"annex-cloud-auth-jwt/router"
	"annex-cloud-auth-jwt/service"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	//db.Table("roles").AutoMigrate(&entity.Role{})
	//db.Table("users").AutoMigrate(&entity.User{})

	//Init Repository
	//which returns the address of userRepository
	//the caller of NewUserRepository gets pointer of UserRepository
	userRepository := repository.NewUserRepository(db)

	//Init Service
	authenticationService := service.NewAuthenticationService(userRepository, validate)

	//Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)

	userController := controller.NewUsersController(userRepository)

	routes := router.NewRouter(userRepository, authenticationController, userController)

	// Create a new Gin router instance
	//gin.Default contains the engine in which request and response can run
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // Add your frontend URL here
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	}))

	// Mount the routes
	//gin.Context = struct contains both httpRequest and Response in Gin Framework which allow users to write their own response also
	router.Any("/*p", func(c *gin.Context) {
		routes.ServeHTTP(c.Writer, c.Request)
	})

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe() //let us connect on tcp address and checks for any incoming request
	helper.ErrorPanic(server_err)
	if err != nil {
        log.Fatal("🚀 Failed to server", err)
    }
}
