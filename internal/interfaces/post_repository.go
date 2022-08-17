package interfaces

import "github.com/yudegaki/apprunner-memoapp/internal/entities"

type PostRepository interface {
	GetDetail(id uint) (*entities.Post, error)
	GetAll() ([]*entities.Post, error)
	Create(post *entities.Post) error
	Update(post *entities.Post) error
	Delete(id uint) error
}
