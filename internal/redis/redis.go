package redis

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type RedisClient struct {
	ctx  context.Context
	conn net.Conn
}

type RedisConn struct {
	Host     string
	Password string
	Port     int
	Db       int
}

func NewRedisClient(conf RedisConn) RedisClient {
	ctx := context.Background()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	if err != nil {
		fmt.Printf("ERR:%v", err)
		os.Exit(1)
	}
	check := func(cmd string) {
		_, err := conn.Write([]byte(cmd))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if conf.Password != "" {
		cmd := fmt.Sprintf("AUTH %s\r\n", conf.Password)
		check(cmd)
	}

	if conf.Db != 0 {
		cmd := fmt.Sprintf("SELECT %d\r\n", conf.Db)
		check(cmd)
	}

	return RedisClient{
		ctx:  ctx,
		conn: conn,
	}
}

func (c RedisClient) commandBytes(cmd string, args ...string) []byte {
	var cmdbuf bytes.Buffer
	fmt.Fprintf(&cmdbuf, "*%d\r\n$%d\r\n%s\r\n", len(args)+1, len(cmd), cmd)
	for _, s := range args {
		fmt.Fprintf(&cmdbuf, "$%d\r\n%s\r\n", len(s), s)
	}
	return cmdbuf.Bytes()
}

func (c *RedisClient) Run() {
	for {
		fmt.Printf(">>")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		cmdSlice := strings.Split(line, " ")
		var cmdSliceRemoveSpace []string
		for _, e := range cmdSlice {
			if len(e) < 1 {
				continue
			}
			e = strings.TrimSpace(e)
			if e == "exit" {
				fmt.Printf("Bye...")
				os.Exit(0)
			}
			cmdSliceRemoveSpace = append(cmdSliceRemoveSpace, e)
		}
		msg := c.commandBytes(cmdSliceRemoveSpace[0], cmdSliceRemoveSpace[1:]...)

		_, err := c.conn.Write(msg)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		r := bufio.NewReader(c.conn)
		res, _ := r.ReadString('\n')
		switch res[0] {
		case '*':
			res = strings.Trim(res, "\r\n")
			res = strings.TrimSpace(res)
			lines, err := strconv.Atoi(res[1:])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for i := 0; i < lines; i++ {
				_, _ = r.ReadString('\n')
				res, _ = r.ReadString('\n')
				fmt.Printf(res)
			}
		case '$':
			res, _ = r.ReadString('\n')
			fmt.Printf(res)
		case '+':
			fmt.Printf(res)
		case ':':
			fmt.Printf(res[1:])
		default:
			fmt.Printf(res)
		}
	}
}
