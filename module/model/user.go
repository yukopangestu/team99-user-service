package model

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primary_key;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type GetListingRequest struct {
	PageNum  int `json:"page_num" query:"page_num"`
	PageSize int `json:"page_size" query:"page_size"`
	UserId   int `json:"user_id" query:"user_id"`
}

type PostListingRequest struct {
	UserId      int    `json:"user_id" form:"user_id" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	ListingType string `json:"listing_type" form:"listing_type" validate:"required"`
}
