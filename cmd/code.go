/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/laoningmeng/fusion/internal/code"
	"github.com/spf13/cobra"
)

var author string
var startDate string
var endDate string

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "统计代码量的工具",
	Long:  `统计代码量的工具`,
	Run: func(cmd *cobra.Command, args []string) {
		c := code.Code{}
		c.Run(author, startDate, endDate)
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	codeCmd.Flags().StringVarP(&author, "author", "a", "", "git用户名字")
	codeCmd.Flags().StringVarP(&startDate, "start", "s", "", "根据时间区间统计代码量-开始时间")
	codeCmd.Flags().StringVarP(&endDate, "end", "e", "", "根据时间区间统计代码量-结束时间")

}
