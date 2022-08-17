package usecases

import "github.com/yudegaki/apprunner-memoapp/internal/interfaces"

type DeletePostUsecase struct {
	repository interfaces.PostRepository
}

func NewDeletePostUsecase(r interfaces.PostRepository) *DeletePostUsecase {
	return &DeletePostUsecase{
		repository: r,
	}
}

func (u *DeletePostUsecase) Execute(id uint) error {
	return u.repository.Delete(id)
}
