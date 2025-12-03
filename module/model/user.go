package model

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primary_key;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type GetUserRequest struct {
	PageNum  int `json:"page_num" query:"page_num"`
	PageSize int `json:"page_size" query:"page_size"`
}

type PostUserRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}
