package bpm

import (
	"fmt"
	"github.com/laoningmeng/fusion/internal/bpm/app/model"
	router2 "github.com/laoningmeng/fusion/internal/bpm/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

type BpmServer struct {
	port int32
}

func NewBpmServer(port int32) *BpmServer {
	return &BpmServer{port: port}
}
func (bpm *BpmServer) boot() {
	path, _ := os.UserHomeDir()
	filePath := path + "/.bpm/bpm.db"
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		os.Mkdir(path+"/.bpm", os.ModePerm)
		_, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		db.AutoMigrate(&model.BpmSettings{})
	}
}

func (bpm BpmServer) Start() {
	router := router2.NewRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", bpm.port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
