package repository

import (
	"github.com/rudikurniawan99/go-api/src/model"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(database *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: database,
	}
}

func (r *AuthorRepository) Create(author *model.Author) error {
	err := r.db.Create(author).Error

	if err != nil {
		return err
	}
	return nil
}
