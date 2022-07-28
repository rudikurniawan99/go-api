package repository

import (
	"github.com/rudikurniawan99/go-api/src/model"
	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) model.BlogRepository {
	return &blogRepository{db}
}

func (r *blogRepository) Create(blog *model.Blog) error {
	err := r.db.Create(&blog).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *blogRepository) GetAll(blog *[]model.Blog) error {
	err := r.db.Find(blog, 10).Error

	if err != nil {
		return err
	}
	return nil
}
