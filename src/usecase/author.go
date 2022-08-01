package usecase

import (
	"github.com/rudikurniawan99/go-api/src/model"
)

type authorUsecase struct {
	r model.AuthorRepository
}

func NewAuthorUsecase(ar model.AuthorRepository) model.AuthorUsecase {
	return &authorUsecase{
		r: ar,
	}

}

func (u *authorUsecase) CreateAuthor(author *model.Author) error {
	err := u.r.Create(author)

	if err != nil {
		return err
	}

	return nil
}
