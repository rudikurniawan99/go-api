package usecase

import "github.com/rudikurniawan99/go-api/src/model"

type blogUsecase struct {
	r model.BlogRepository
}

func NewBlogUsecase(br model.BlogRepository) model.BlogUsecase {
	return &blogUsecase{
		r: br,
	}
}

func (u *blogUsecase) Create(blog *model.Blog) error {

	err := u.r.Create(blog)

	if err != nil {
		return err
	}

	return nil
}

func (u *blogUsecase) FetchAll(blog *[]model.Blog) error {
	err := u.r.GetAll(blog)

	if err != nil {
		return err
	}

	return nil
}
