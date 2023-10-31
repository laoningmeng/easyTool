package bpm

import (
	"fmt"
	"github.com/laoningmeng/fusion/internal/bpm/app/model"
	"github.com/laoningmeng/fusion/internal/bpm/global"
	router2 "github.com/laoningmeng/fusion/internal/bpm/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net"
	"net/http"
	"os"
	"time"
)

type BpmServer struct {
	port int
}

func init() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	global.DB = db
}

func NewBpmServer(port int) *BpmServer {
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
		global.DB.AutoMigrate(&model.BpmSettings{})
	}
}

func (bpm BpmServer) Start() {
	router := router2.NewRouter()
	var port int
	if bpm.port == 0 {
		port, _ = bpm.getPort()
	} else {
		port = bpm.port
	}
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Port:", port)
	s.ListenAndServe()
}

func (s *BpmServer) getPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
