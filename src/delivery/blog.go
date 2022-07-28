package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api/helper"
	"github.com/rudikurniawan99/go-api/src/model"
)

type blogDelivery struct {
	u model.BlogUsecase
}

func NewDelivery(bu model.BlogUsecase) model.BlogDelivery {
	return &blogDelivery{
		u: bu,
	}
}

func (d *blogDelivery) Mount(group *gin.RouterGroup) {
	group.GET("/", d.GetAllBlogHandler)
	group.POST("/", d.CreateHandler)
}

func (d *blogDelivery) GetAllBlogHandler(c *gin.Context) {
	blogs := &[]model.Blog{}

	err := d.u.FetchAll(blogs)

	if err != nil {
		helper.JsonError(c, err)
		return
	}

	helper.JsonSUCCESS(c, blogs)
}

func (d *blogDelivery) CreateHandler(c *gin.Context) {
	title := c.PostForm("title")
	desc := c.PostForm("description")

	blog := &model.Blog{
		Title:       title,
		Description: desc,
	}

	err := d.u.Create(blog)

	if err != nil {
		helper.JsonError(c, err)
	}

	helper.JsonSUCCESS(c, blog)
}
