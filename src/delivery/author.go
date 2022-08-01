package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api/helper"
	"github.com/rudikurniawan99/go-api/src/model"
)

type authorDelivery struct {
	authorUsecase model.AuthorUsecase
}

func (a *authorDelivery) Mount(group *gin.RouterGroup) {
	group.POST("/", a.CreateAuthor)
}

func (a *authorDelivery) CreateAuthor(c *gin.Context) {
	req := model.AuthorRequest{}
	c.Bind(&req)

	author := &model.Author{
		Name: req.Name,
		Age:  req.Age,
	}

	err := a.authorUsecase.CreateAuthor(author)
	if err != nil {
		helper.JsonError(c, err)
		return
	}

	helper.JsonSUCCESS(c, author)
}

func NewAuthorDelivery(au model.AuthorUsecase) model.AuthorDelivery {
	return &authorDelivery{
		authorUsecase: au,
	}
}
