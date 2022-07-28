package model

import "gorm.io/gorm"

type (
	Blog struct {
		gorm.Model
		Title       string `gorm:"type:varchar(250)"`
		Description string `gorm:"type:text"`
	}

	BlogRepository interface {
		Create(blog *Blog) error
		GetAll(blog *[]Blog) error
	}
)
