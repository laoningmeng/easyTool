/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/laoningmeng/fusion/internal/tools"
	"github.com/spf13/cobra"
)

var port int

// toolsCmd represents the tools command
var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "开发工具集",
	Long:  `开发工具集`,
	Run: func(cmd *cobra.Command, args []string) {
		s := tools.NewServer()
		s.Run(port)
	},
}

func init() {
	rootCmd.AddCommand(toolsCmd)
	toolsCmd.Flags().IntVarP(&port, "port", "p", 0, "端口号")
}
