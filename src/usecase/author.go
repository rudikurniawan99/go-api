package usecase

import (
	"github.com/rudikurniawan99/go-api/src/model"
	"github.com/rudikurniawan99/go-api/src/repository"
)

type AuthorUsecase struct {
	u *repository.AuthorRepository
}

func NewAuthorUsecase(ar *repository.AuthorRepository) *AuthorUsecase {
	return &AuthorUsecase{
		u: ar,
	}
}

func (u *AuthorUsecase) CreateData(author *model.Author) error {
	err := u.u.Create(author)

	if err != nil {
		return err
	}

	return nil
}
