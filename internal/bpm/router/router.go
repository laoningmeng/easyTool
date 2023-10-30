package router

import (
	"github.com/gin-gonic/gin"
	bpm2 "github.com/laoningmeng/fusion/internal/bpm/app/controller/bpm"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	setting := bpm2.Settings{}
	{
		router.POST("/add", setting.Add)
		router.POST("/update", setting.Add)
		router.GET("/del", setting.Add)
		router.GET("/list", setting.Add)
	}
	return router
}
