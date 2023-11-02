package bpm

import (
	"fmt"
	model2 "github.com/laoningmeng/fusion/internal/bpm/app/model"
	"github.com/laoningmeng/fusion/internal/bpm/global"
	router2 "github.com/laoningmeng/fusion/internal/bpm/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net"
	"net/http"
	"os"
	"time"
)

var homeDir, _ = os.UserHomeDir()
var db_path = homeDir + "/.bpm/bpm.db"

type BpmServer struct {
	port int
}

func init() {
	_, err := os.Stat(db_path)
	if os.IsNotExist(err) {
		db, _ := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
		db.AutoMigrate(model2.BpmSettings{})
		global.DB = db
	} else {
		db, _ := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
		global.DB = db
	}
}

func NewBpmServer(port int) *BpmServer {
	return &BpmServer{port: port}
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
