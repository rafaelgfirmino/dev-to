package service

import (
	"github.com/rafaelgfirmino/dev_to/using_tilt/internal/domain"
	"github.com/rafaelgfirmino/dev_to/using_tilt/internal/port"
)

type UserService struct {
	userRepository port.UserDbRepository
}

func NewUserService(userRepository port.UserDbRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
func (s *UserService) GetUserById(id string) (*domain.User, error) {
	user, err := (s.userRepository).FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *UserService) GetUserByEmail(email string) (*domain.User, error) {
	user, err := (s.userRepository).FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
