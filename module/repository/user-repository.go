package repository

import (
	"team99_listing_service/module/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(id string) (model.User, error)
	CreateUser(listing model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
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
