package service

import (
	"team99_listing_service/module/model"
	"team99_listing_service/module/repository"
)

type UserServiceInterface interface {
	GetAllUser(request model.GetUserRequest) ([]model.User, error)
	GetUserById(id string) (model.User, error)
	PostUser(request model.PostUserRequest) (model.User, error)
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserServiceInterface {
	return &userService{UserRepository: userRepository}
}

func (s userService) GetAllUser(request model.GetUserRequest) ([]model.User, error) {
	data, err := s.UserRepository.GetUser(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s userService) GetUserById(id string) (model.User, error) {
	data, err := s.UserRepository.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}

	return data, nil
}

func (s userService) PostUser(request model.PostUserRequest) (model.User, error) {
	var data model.User

	data = model.User{
		Name: request.Name,
	}

	result, err := s.UserRepository.CreateUser(data)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}
