package tools

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
)

//go:embed *
var Resource embed.FS

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port int) {
	if port == 0 {
		var err error
		port, err = s.getPort()
		if err != nil {
			os.Exit(1)
		}
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.StaticFS("/", http.FS(Resource))
	go func() {
		r.Run(":" + strconv.Itoa(port))
	}()
	fmt.Printf("\u001B[33;1m Tools Port:%d \n", port)
	s.Open("http://localhost:" + strconv.Itoa(port))
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}

func (s *Server) getPort() (int, error) {
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
func (s *Server) Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		os.Exit(1)
	}
	cmd := exec.Command(run, uri)
	return cmd.Start()
}
