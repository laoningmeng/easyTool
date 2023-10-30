package bpm

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Settings struct {
}

type SettingItem struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	FormKey  string `json:"form_key,omitempty"`
	AppId    string `json:"app_id,omitempty"`
	DataCode string `json:"data_code,omitempty"`
	Fields   string `json:"fields,omitempty"`
}

func (s *Settings) Add(c *gin.Context) {

}

func (s *Settings) Update(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (s *Settings) List(c *gin.Context) {

}

func (s *Settings) Del(c *gin.Context) {

}
