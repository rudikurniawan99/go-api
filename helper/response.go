package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	jsonSuccess struct {
		Message string      `json:"message"`
		Status  string      `json:"status"`
		Data    interface{} `json:"data"`
	}
	jsonError struct {
		Message string      `json:"message"`
		Status  string      `json:"status"`
		Data    interface{} `json:"data,omitempty"`
	}
)

func JsonSUCCESS(c *gin.Context, data interface{}) {
	res := jsonSuccess{
		Message: "success",
		Status:  "success",
		Data:    data,
	}

	c.JSON(http.StatusOK, res)
}

func JsonError(c *gin.Context, err error) {
	res := jsonError{
		Message: err.Error(),
		Status:  "failed",
		Data:    nil,
	}

	c.JSON(http.StatusBadRequest, res)
}

func JsonSuccessCustomMessage(c *gin.Context, msg string, data interface{}) {
	res := jsonSuccess{
		Message: msg,
		Status:  "success",
		Data:    data,
	}

	c.JSON(http.StatusOK, res)
}
