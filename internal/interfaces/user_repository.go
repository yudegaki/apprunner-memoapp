package interfaces

import "github.com/yudegaki/apprunner-memoapp/internal/entities"

type UserRepository interface {
	Get() ([]*entities.User, error)
	GetDetail(id uint) (*entities.User, error)
}
