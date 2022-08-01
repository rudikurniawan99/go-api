package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Author struct {
		gorm.Model
		Name string `gorm:"varchar(50)"`
		Age  int
	}

	AuthorRequest struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	AuthorUsecase interface {
		CreateAuthor(author *Author) error
	}

	AuthorRepository interface {
		Create(author *Author) error
	}

	AuthorDelivery interface {
		Mount(group *gin.RouterGroup)
	}
)
