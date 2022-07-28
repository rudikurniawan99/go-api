package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	d.u.FetchAll(blogs)

	c.JSON(http.StatusOK, gin.H{
		"message": true,
		"data":    blogs,
	})
}

func (d *blogDelivery) CreateHandler(c *gin.Context) {
	title := c.PostForm("title")
	desc := c.PostForm("description")

	blog := &model.Blog{
		Title:       title,
		Description: desc,
	}

	d.u.Create(blog)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "success create blog",
		"data":    blog,
	})
}
