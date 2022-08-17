package usecases

import (
	"github.com/yudegaki/apprunner-memoapp/internal/entities"
	"github.com/yudegaki/apprunner-memoapp/internal/interfaces"
)

type GetUsersUsecase struct {
	repository interfaces.UserRepository
}

func NewGetUsersUsecase(r interfaces.UserRepository) *GetUsersUsecase {
	return &GetUsersUsecase{
		repository: r,
	}
}

func (u *GetUsersUsecase) Execute() ([]*entities.User, error) {
	return u.repository.Get()
}
