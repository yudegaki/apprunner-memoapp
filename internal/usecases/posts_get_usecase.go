package usecases

import (
	"github.com/yudegaki/apprunner-memoapp/internal/entities"
	"github.com/yudegaki/apprunner-memoapp/internal/interfaces"
)

type GetPostsUsecase struct {
	repository interfaces.PostRepository
}

func NewGetPostsUsecase(r interfaces.PostRepository) *GetPostsUsecase {
	return &GetPostsUsecase{
		repository: r,
	}
}

func (u *GetPostsUsecase) Execute() ([]*entities.Post, error) {
	return u.repository.GetAll()
}
