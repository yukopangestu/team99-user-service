package repository

import (
	"team99_listing_service/module/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetListing(request model.GetListingRequest) ([]model.User, error)
	CreateListing(listing model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) GetListing(request model.GetListingRequest) ([]model.User, error) {
	var list []model.User
	query := r.db.Model(&model.User{})
	if request.UserId != 0 {
		query.Where("user_id = ?", request.UserId)
	}

	if request.PageNum == 0 {
		request.PageNum = 1
	}

	if request.PageSize == 0 {
		request.PageSize = 10
	}

	query.Limit(request.PageSize).Offset((request.PageNum - 1) * request.PageSize)
	err := query.Find(&list).Error
	return list, err
}

func (r userRepository) CreateListing(listing model.User) (model.User, error) {
	err := r.db.Create(&listing).Error

	return listing, err
}
