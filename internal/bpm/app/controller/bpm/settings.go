package bpm

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	model2 "github.com/laoningmeng/fusion/internal/bpm/app/model"
	"github.com/laoningmeng/fusion/internal/bpm/global"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Settings struct {
}

type SettingRequest struct {
	Id       string  `json:"id"`
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
type PageInfo struct {
	Page  int `json:"curr"`
	Limit int `json:"limit"`
}
type Query struct {
	Page     PageInfo `json:"page"`
	Name     string   `json:"name"`
	DataCode string   `json:"data_code"`
	AppId    string   `json:"app_id"`
}

type SettingList struct {
	Total int        `json:"total"`
	Code  int        `json:"code"`
	Data  []Settings `json:"data"`
}
type SettingItem struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	FormKey   string    `json:"form_key"`
	AppId     string    `json:"app_id"`
	DataCode  string    `json:"data_code"`
	Fields    []Field   `json:"fields"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Settings) Add(c *gin.Context) {
	var param SettingRequest
	if err := c.ShouldBind(&param); err == nil {
		fieldByte, _ := json.Marshal(param.Fields)
		var checkSetting model2.BpmSettings
		result := global.DB.Where("data_code=?", param.DataCode).First(&checkSetting)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			data := model2.BpmSettings{
				Uuid:      uuid.New().String(),
				Name:      param.Name,
				FormKey:   param.FormKey,
				AppId:     param.AppId,
				DataCode:  param.DataCode,
				Fields:    string(fieldByte),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			global.DB.Create(data)
			httpOk(c, "")
			return
		}
		httpFail(c, "already exists")
		return
	}
	httpFail(c, "miss param")
	return
}

func (s *Settings) Update(c *gin.Context) {
	var param SettingRequest
	if err := c.ShouldBind(&param); err == nil {

		result := global.DB.Where("data_code=?", param.Id).First(&model2.BpmSettings{})
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			httpFail(c, "not found settings")
			return
		}
		fieldByte, _ := json.Marshal(param.Fields)
		model := model2.BpmSettings{
			Uuid:      param.Id,
			Name:      param.Name,
			FormKey:   param.FormKey,
			AppId:     param.AppId,
			DataCode:  param.DataCode,
			Fields:    string(fieldByte),
			UpdatedAt: time.Now(),
		}
		tx := global.DB.Model(model2.BpmSettings{}).Where("uuid=?", param.Id).Updates(&model)
		if tx.Error != nil {
			httpFail(c, tx.Error.Error())
			return
		}
		httpOk(c, "success")
		return
	}
	httpFail(c, "miss param")
	return
}

func (s *Settings) List(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	limitStr := c.DefaultQuery("limit", "50")
	limit, _ := strconv.Atoi(limitStr)
	name := c.DefaultQuery("name", "")
	appId := c.DefaultQuery("app_id", "")
	dataCode := c.DefaultQuery("data_code", "")

	start := (page - 1) * limit
	var data []model2.BpmSettings
	global.DB.Where(model2.BpmSettings{
		Name:     name,
		AppId:    appId,
		DataCode: dataCode,
	}).Offset(start).Limit(100).Find(&data)
	var result []SettingItem
	for _, value := range data {
		var field []Field
		json.Unmarshal([]byte(value.Fields), &field)
		result = append(result, SettingItem{
			Id:        value.Uuid,
			Name:      value.Name,
			FormKey:   value.FormKey,
			AppId:     value.AppId,
			DataCode:  value.DataCode,
			Fields:    field,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}
	httpOk(c, result)
	return
}

func (s *Settings) Del(c *gin.Context) {
	var param SettingRequest
	if err := c.ShouldBind(&param); err == nil {

		result := global.DB.Where("uuid=?", param.Id).First(&model2.BpmSettings{})
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			httpFail(c, "not found settings")
			return
		}

		tx := global.DB.Model(model2.BpmSettings{}).Where("uuid=?", param.Id).Delete(model2.BpmSettings{})
		if tx.Error != nil {
			httpFail(c, tx.Error.Error())
			return
		}
		httpOk(c, "success")
		return
	}
	httpFail(c, "miss param")
	return
}
