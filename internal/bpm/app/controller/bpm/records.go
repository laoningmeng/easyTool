package bpm

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Records struct {
}

func (r *Records) Add(c *gin.Context) {

}

func (r *Records) Update(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (r *Records) List(c *gin.Context) {

}

func (r *Records) Del(c *gin.Context) {

}

func (r *Records) Send(c *gin.Context) {

}
