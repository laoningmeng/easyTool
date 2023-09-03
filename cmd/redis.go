/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/laoningmeng/fusion/internal/redis"
	"github.com/spf13/cobra"
)

var redisHost string
var redisPort int
var redisPassword string
var redisDb int

// redisCmd represents the redis command
var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "redis client",
	Long:  `一个连接redis 的命令行工具`,
	Run: func(cmd *cobra.Command, args []string) {
		c := redis.NewRedisClient(redis.RedisConn{
			Host:     redisHost,
			Password: redisPassword,
			Port:     redisPort,
			Db:       redisDb,
		})
		c.Run()
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
	redisCmd.Flags().StringVarP(&redisHost, "host", "", "127.0.0.1", "redis host")
	redisCmd.Flags().StringVarP(&redisPassword, "pass", "", "", "redis pass")
	redisCmd.Flags().IntVarP(&redisDb, "db", "", 0, "第几个db，默认0")
	redisCmd.Flags().IntVarP(&redisPort, "port", "p", 6379, "每次插入数据表的数量")
}
