package service

import (
	"wire-hzw/internal/model"
	"wire-hzw/internal/repository"
)

type UserService interface {
	GetUserById(id int64) (*model.User, error)
}

type userService struct {
	*Service
	userRepository repository.UserRepository
}

func NewUserService(service *Service, userRepository repository.UserRepository) UserService {
	return &userService{
		Service:        service,
		userRepository: userRepository,
	}
}

func (s *userService) GetUserById(id int64) (*model.User, error) {
	return s.userRepository.FirstById(id)
}
