package code

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Code struct {
}

func (c *Code) Run(author, start, end string) {
	c.getCodeTotal(author, start, end)
}

func (c *Code) getGitAuthor() (string, error) {
	cmd := exec.Command("git", "config", "user.name")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("ERR:", err)
		os.Exit(1)
	}
	name := string(output)
	name = strings.Trim(name, "\n")
	return name, nil
}

// 根据用户名时间段统计

func (c *Code) getCodeTotal(name, start, end string) {
	if name == "" {
		var err error
		name, err = c.getGitAuthor()
		if err != nil {
			fmt.Println("获取git用户信息失败")
			os.Exit(1)
		}
	}
	currentTime := time.Now()
	if start == "" {
		start = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Format("2006-01-02 15:04:05")
	}
	if end == "" {
		end = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).Format("2006-01-02 15:04:05")
	}

	cmd := exec.Command("git", "log", "--author="+name, "--since="+start, "--until="+end, "--pretty=tformat:", "--numstat")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(output), "\n")
	var add int
	var remove int
	for _, line := range lines {
		reg, err := regexp.Compile(`(\d+)\s`)
		if err != nil {
			panic(err)
		}
		res := reg.FindAllString(line, -1)
		if len(res) < 2 {
			continue
		}
		convert := func(res string) int {
			res = strings.Trim(res, "\t")
			number, err := strconv.Atoi(res)
			if err != nil {
				panic(err)
			}
			return number
		}
		add = add + convert(res[0])
		remove = remove + convert(res[1])
	}
	fmt.Printf("\u001B[33;1m 用户:%s \u001B[34;1m 新增:%d 行, \u001B[31;1m 删除:%d 行, \u001B[36;1m 共计: %d 行 \u001B[0m \n", name, add, remove, add-remove)

}
