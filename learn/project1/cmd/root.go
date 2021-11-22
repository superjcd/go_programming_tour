package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{}
	str     string
	mode    int8
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
