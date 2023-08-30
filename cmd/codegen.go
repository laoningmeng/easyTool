/*
*
摘取自 https://github.com/marmotedu/iam/tree/master/tools/codegen
*/
package cmd

import (
	"github.com/laoningmeng/fusion/internal/codegen"
	"github.com/spf13/cobra"
)

var (
	TypeName   string
	Output     string
	Trimprefix string
	BuildTags  string
	Doc        bool
)

// codegenCmd represents the codegen command
var codegenCmd = &cobra.Command{
	Use:   "codegen",
	Short: "异常码生成工具",
	Long:  `异常码生成工具`,
	Run: func(cmd *cobra.Command, args []string) {
		c := &codegen.CodeGen{
			TypeName:   TypeName,
			Output:     Output,
			Trimprefix: Trimprefix,
			BuildTags:  BuildTags,
			Doc:        Doc,
		}
		c.Run()
	},
}

func init() {
	rootCmd.AddCommand(codegenCmd)
	codegenCmd.Flags().StringVarP(&TypeName, "type", "t", "", "comma-separated list of type names; must be set")
	codegenCmd.Flags().StringVarP(&Output, "output", "o", "", "output file name; default srcdir/<type>_string.go")
	codegenCmd.Flags().StringVarP(&Trimprefix, "trimprefix", "", "", "comma-separated list of build tags to apply")
	codegenCmd.Flags().StringVarP(&BuildTags, "tags", "", "", "trim the `prefix` from the generated constant names")
	codegenCmd.Flags().BoolVarP(&Doc, "doc", "d", false, "if true only generate error code documentation in markdown format")
}
