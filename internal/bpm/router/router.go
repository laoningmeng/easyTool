package router

import (
	"github.com/gin-gonic/gin"
	bpm2 "github.com/laoningmeng/fusion/internal/bpm/app/controller/bpm"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	setting := bpm2.Settings{}
	record := bpm2.Records{}
	{
		router.POST("/bpm/setting/add", setting.Add)
		router.POST("/bpm/setting/update", setting.Update)
		router.GET("/bpm/setting/del", setting.Del)
		router.GET("/bpm/setting/list", setting.List)

		router.POST("/bpm/record/add", record.Add)
		router.POST("/bpm/record/update", record.Update)
		router.POST("/bpm/record/send", record.Update)
		router.GET("/bpm/record/del", record.Del)
		router.GET("/bpm/record/list", record.Add)
	}
	return router
}
