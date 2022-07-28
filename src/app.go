package src

import (
	"github.com/gin-gonic/gin"
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
	s.httpServer.GET(":name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "success",
			"data":    name,
		})
	})

	s.httpServer.Run(":8082")
}
