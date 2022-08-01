package repository

import (
	"github.com/rudikurniawan99/go-api/src/model"
	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(database *gorm.DB) model.AuthorRepository {
	return &authorRepository{
		db: database,
	}
}

func (r *authorRepository) Create(author *model.Author) error {
	err := r.db.Create(author).Error

	if err != nil {
		return err
	}
	return nil
}
