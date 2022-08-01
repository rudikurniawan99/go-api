package src

import (
	"github.com/gin-gonic/gin"
	"github.com/rudikurniawan99/go-api/registry"
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
	blogGroup := s.httpServer.Group("blog")
	blogDelivery.Mount(blogGroup)

	authorDelivery := registry.AuthorDelivery(s.database)
	authorGroup := s.httpServer.Group("author")
	authorDelivery.Mount(authorGroup)

	s.httpServer.Run(":8082")
}
