package registry

import (
	"github.com/rudikurniawan99/go-api/src/delivery"
	"github.com/rudikurniawan99/go-api/src/model"
	"github.com/rudikurniawan99/go-api/src/repository"
	"github.com/rudikurniawan99/go-api/src/usecase"
	"gorm.io/gorm"
)

func BlogDelivery(db *gorm.DB) model.BlogDelivery {
	blogRepository := repository.NewRepository(db)
	blogUsecase := usecase.NewBlogUsecase(blogRepository)
	blogDelivery := delivery.NewDelivery(blogUsecase)
	return blogDelivery
}

func AuthorDelivery(db *gorm.DB) model.AuthorDelivery {

	authorRepository := repository.NewAuthorRepository(db)
	authorUsecase := usecase.NewAuthorUsecase(authorRepository)
	authorDelivery := delivery.NewAuthorDelivery(authorUsecase)

	return authorDelivery
}
