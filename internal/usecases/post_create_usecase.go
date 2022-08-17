package usecases

import (
	"github.com/yudegaki/apprunner-memoapp/internal/entities"
	"github.com/yudegaki/apprunner-memoapp/internal/interfaces"
)

type CreatePostUsecase struct {
	repository interfaces.PostRepository
}

func NewCreatePostUsecase(r interfaces.PostRepository) *CreatePostUsecase {
	return &CreatePostUsecase{
		repository: r,
	}
}

func (u *CreatePostUsecase) Execute(post entities.Post) error {
	return u.repository.Create(&post)
}
