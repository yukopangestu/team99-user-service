package service

import (
	"team99_listing_service/module/model"
	"team99_listing_service/module/repository"
)

type UserServiceInterface interface {
	GetUserById(id string) (model.User, error)
	PostListing(request model.PostListingRequest) (model.User, error)
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserServiceInterface {
	return &userService{UserRepository: userRepository}
}

func (s userService) GetUserById(id string) (model.User, error) {
	data, err := s.UserRepository.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}

	return data, nil
}

func (s userService) PostListing(request model.PostListingRequest) (model.User, error) {
	var data model.User

	//TODO: Validate user_id with the other services

	data = model.User{
		UserId:      request.UserId,
		Price:       request.Price,
		ListingType: request.ListingType,
	}

	result, err := s.UserRepository.CreateUser(data)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}
