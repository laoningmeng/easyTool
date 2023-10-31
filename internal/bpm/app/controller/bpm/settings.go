package bpm

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Settings struct {
}

type SettingRequest struct {
	Name     string  `json:"name"`
	FormKey  string  `json:"form_key"`
	AppId    string  `json:"app_id"`
	DataCode string  `json:"data_code"`
	Fields   []Field `json:"fields"`
}

type Field struct {
	Title string `json:"title"`
	Name  string `json:"name"`
}

func (s *Settings) Add(c *gin.Context) {
	var param SettingRequest
	if errA := c.ShouldBind(&param); errA == nil {
		fmt.Println(param)
	}
}

func (s *Settings) Update(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (s *Settings) List(c *gin.Context) {
	//
}

func (s *Settings) Del(c *gin.Context) {

}
