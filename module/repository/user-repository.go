package repository

import (
	"team99_listing_service/module/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(request model.GetUserRequest) ([]model.User, error)
	GetUserById(id string) (model.User, error)
	CreateUser(listing model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) GetUser(request model.GetUserRequest) ([]model.User, error) {
	var user []model.User
	query := r.db.Model(&model.User{})
	if request.PageNum == 0 {
		request.PageNum = 1
	}

	if request.PageSize == 0 {
		request.PageSize = 10
	}

	query.Limit(request.PageSize).Offset((request.PageNum - 1) * request.PageSize)
	err := query.Find(&user).Error
	return user, err
}

func (r userRepository) GetUserById(id string) (model.User, error) {
	var user model.User
	err := r.db.Where("user_id = ?", id).Find(&user).Error
	return user, err
}

func (r userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
