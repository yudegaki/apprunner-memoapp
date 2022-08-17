package usecases

import (
	"github.com/yudegaki/apprunner-memoapp/internal/entities"
	"github.com/yudegaki/apprunner-memoapp/internal/interfaces"
)

type GetPostUsecase struct {
	repository interfaces.PostRepository
}

func NewGetPostUsecase(r interfaces.PostRepository) *GetPostUsecase {
	return &GetPostUsecase{
		repository: r,
	}
}

func (u *GetPostUsecase) Execute(id uint) (*entities.Post, error) {
	return u.repository.GetDetail(id)
}
