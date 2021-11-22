package cmd

import (
	"log"
	"project1/internal/word"
	"strings"

	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
)

var desc = strings.Join([]string{
	"1 转化为大写",
	"2 转化为小写",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "文本格式转化",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果: %s", content)
	},
}
