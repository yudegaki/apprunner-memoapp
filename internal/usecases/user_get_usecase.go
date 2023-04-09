package usecases

import (
	"github.com/yudegaki/apprunner-memoapp/internal/entities"
	"github.com/yudegaki/apprunner-memoapp/internal/interfaces"
)

type GetUserUsecase struct {
	repository interfaces.UserRepository
}

func NewGetUserUsecase(r interfaces.UserRepository) *GetUserUsecase {
	return &GetUserUsecase{
		repository: r,
	}
}

func (u *GetUserUsecase) Execute(id uint) (*entities.User, error) {
	return u.repository.GetDetail(id)
}
