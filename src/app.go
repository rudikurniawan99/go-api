package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api/src/model"
	"github.com/rudikurniawan99/go-api/src/repository"
	"github.com/rudikurniawan99/go-api/src/utils/db"
	"gorm.io/gorm"
)

type server struct {
	httpServer *gin.Engine
	database   *gorm.DB
}

func Init() *server {
	s := &server{
		httpServer: gin.Default(),
		database:   db.Connect(),
	}

	return s
}

func (s *server) Run() {
	s.httpServer.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "connection test success",
		})
	})

	// -
	blogRepository := repository.NewRepository(s.database)

	s.httpServer.POST("blog", func(c *gin.Context) {
		title := c.PostForm("title")
		description := c.PostForm("description")

		blog := &model.Blog{
			Title:       title,
			Description: description,
		}

		blogRepository.Create(blog)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "get all data",
			"data":    blog,
		})
	})

	s.httpServer.GET("blog", func(c *gin.Context) {
		blogs := &[]model.Blog{}

		blogRepository.GetAll(blogs)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "get all data",
			"data":    blogs,
		})
	})

	s.httpServer.Run(":8082")
}
