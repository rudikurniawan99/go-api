package src

import (
	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api/src/delivery"
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
	blogRepository := repository.NewRepository(s.database)
	blogUsecase := usecase.NewBlogUsecase(blogRepository)
	blogDelivery := delivery.NewDelivery(blogUsecase)
	authGroup := s.httpServer.Group("blog")
	blogDelivery.Mount(authGroup)

	s.httpServer.Run(":8082")
}
