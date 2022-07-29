package src

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api/helper"
	"github.com/rudikurniawan99/go-api/registry"
	"github.com/rudikurniawan99/go-api/src/model"
	"github.com/rudikurniawan99/go-api/src/repository"
	"github.com/rudikurniawan99/go-api/src/usecase"
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
	blogDelivery := registry.BlogDelivery(s.database)
	authGroup := s.httpServer.Group("blog")
	blogDelivery.Mount(authGroup)

	authorRepository := repository.NewAuthorRepository(s.database)
	authorUsecase := usecase.NewAuthorUsecase(authorRepository)

	s.httpServer.POST("author", func(c *gin.Context) {
		name := c.PostForm("name")
		age, _ := strconv.Atoi(c.PostForm("age"))

		author := &model.Author{
			Name: name,
			Age:  age,
		}

		err := authorUsecase.CreateData(author)

		if err != nil {
			helper.JsonError(c, err)
			return
		}

		helper.JsonSUCCESS(c, author)
	})

	s.httpServer.Run(":8082")
}
