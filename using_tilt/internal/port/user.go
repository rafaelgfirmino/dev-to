package port

import "github.com/rafaelgfirmino/dev_to/using_tilt/internal/domain"

type UserDbRepository interface {
	FindUserById(id string) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
}

type UserRepository interface {
	GetUserById(id string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
