package cmd

import (
	"log"
	"strings"

	"github.com/bluesky6529/go_tour/internal/word"

	"github.com/spf13/cobra"
)

const (
	ModeUpper                      = iota + 1 // 全轉大寫
	ModeLower                                 // 全轉小寫
	ModeUnderscoreToUpperCamelCase            // 底線轉大寫駝峰單字
	ModeUnderscoreToLowerCamelCase            // 底線轉小寫駝峰單字
	ModeCamelCaseToUnderscore                 // 駝峰單字轉底線單字
)

var str string
var mode int8
var desc = strings.Join([]string{
	"該子命令支援各種單字格是轉換，模式如下：",
	"1：全轉大寫",
	"2：全轉小寫",
	"3：底線轉大寫駝峰單字",
	"4：底線轉小寫駝峰單字",
	"5：駝峰單字轉底線單字",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "單字格式轉換",
	Long:  desc, //help string 上面有定義
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输出结果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "樹入單字內容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "輸入單字轉換模式")
}
