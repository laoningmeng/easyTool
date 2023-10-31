/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/laoningmeng/fusion/internal/bpm"
	"github.com/spf13/cobra"
)

var bpmPort int

// bpmCmd represents the bpm command
var bpmCmd = &cobra.Command{
	Use:   "bpm",
	Short: "简单的bpm流程工具",
	Long:  `简单的bpm流程工具`,
	Run: func(cmd *cobra.Command, args []string) {
		s := bpm.NewBpmServer(bpmPort)
		s.Start()
	},
}

func init() {
	rootCmd.AddCommand(bpmCmd)
	bpmCmd.Flags().IntVarP(&bpmPort, "port", "", 0, "端口号")
}
