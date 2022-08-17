package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yudegaki/apprunner-memoapp/internal/entities"
	"github.com/yudegaki/apprunner-memoapp/internal/external"
	"github.com/yudegaki/apprunner-memoapp/internal/repositories"
	"github.com/yudegaki/apprunner-memoapp/internal/usecases"
)

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func convertEntityUserToControllerUser(entityUser *entities.User) *UserResponse {
	return &UserResponse{
		ID:   entityUser.GetID(),
		Name: entityUser.GetName(),
	}
}

func convertEntityUsersToControllerUsers(entityUsers []*entities.User) []*UserResponse {
	var controllerUsers []*UserResponse
	for _, entityUser := range entityUsers {
		controllerUsers = append(controllerUsers, convertEntityUserToControllerUser(entityUser))
	}
	return controllerUsers
}

func GetAllUsers(c *gin.Context) {
	repository := repositories.NewUserRepository(external.DB)
	usecase := usecases.NewGetUserUsecase(repository)

	result, err := usecase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp := convertEntityUsersToControllerUsers(result)
	c.JSON(200, resp)
}
