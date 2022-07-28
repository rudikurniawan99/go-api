package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Blog struct {
		gorm.Model
		Title       string `gorm:"type:varchar(250)"`
		Description string `gorm:"type:text"`
	}

	BlogRepository interface {
		Create(blog *Blog) error
		GetAll(blog *[]Blog) error
	}

	BlogUsecase interface {
		Create(blog *Blog) error
		FetchAll(blog *[]Blog) error
	}

	BlogDelivery interface {
		Mount(group *gin.RouterGroup)
	}
)
