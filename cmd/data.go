/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/laoningmeng/fusion/internal/data"
	"github.com/spf13/cobra"
)

var path string
var size int

// dataCmd represents the data command
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "mysql通过配置文件生成测试数据",
	Long:  `数据库插入模拟数据的工具`,
	Run: func(cmd *cobra.Command, args []string) {
		f := data.NewMysqlFaker()
		f.Run(path, size)
	},
}

func init() {
	rootCmd.AddCommand(dataCmd)
	dataCmd.Flags().StringVarP(&path, "conf", "c", "", "指定配置文件路径")
	dataCmd.Flags().IntVarP(&size, "size", "s", 100, "每次插入数据表的数量")
}
