package bpm

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func httpOk(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func httpFail(c *gin.Context, msg string) {
	c.JSON(http.StatusBadGateway, msg)
}
