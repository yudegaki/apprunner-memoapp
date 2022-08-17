package usecases

import (
	"github.com/yudegaki/apprunner-memoapp/internal/entities"
	"github.com/yudegaki/apprunner-memoapp/internal/interfaces"
)

type UpdatePostUsecase struct {
	repository interfaces.PostRepository
}

func NewUpdatePostUsecase(r interfaces.PostRepository) *UpdatePostUsecase {
	return &UpdatePostUsecase{
		repository: r,
	}
}

func (u *UpdatePostUsecase) Execute(post entities.Post) error {
	return u.repository.Update(&post)
}
