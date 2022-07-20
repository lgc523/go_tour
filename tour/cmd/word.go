package cmd

import (
	w "github.com/go-tour/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModeUpper = iota + 1
	Mode2Lower
	ModeUnderscore2UpperCamelcase
	ModeUnderscore2LowerCamelcase
	ModeCamelcase2UnderScope
)

var desc = strings.Join([]string{
	"该命令支持各种单词格式转换，模式如下",
	"1:全部单词转为大写",
	"2:全部单词转为小写",
	"3:下划线单词转为大写驼峰",
	"4:下划线单词转为小写驼峰",
	"5:驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run:   wordTask,
}

func wordTask(*cobra.Command, []string) {
	var content string
	switch mode {
	case ModeUpper:
		content = w.ToUpper(str)
	case Mode2Lower:
		content = w.ToLower(str)
	case ModeUnderscore2UpperCamelcase:
		content = w.UnderScoreToUpperCamelCase(str)
	case ModeUnderscore2LowerCamelcase:
		content = w.UnderScoreToLowerCamelCase(str)
	case ModeCamelcase2UnderScope:
		content = w.CamelCaseToUnderScore(str)
	default:
		log.Fatalf("不支持转换格式: input help see usage")
	}
	log.Printf("convert result:%s", content)
}

var mode int
var str string

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "input content")
	wordCmd.Flags().IntVarP(&mode, "mode", "m", 3, "input mode")
}
