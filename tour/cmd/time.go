package cmd

import (
	"log"
	"time"

	"github.com/go-tour/tour/internal/timer"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间戳格式处理",
	Run:   timeHandler,
}

func timeHandler(cmd *cobra.Command, args []string) {
	nowTime := timer.GetNowTime()
	log.Printf("北京时间: %s,%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.UnixMilli())
}

var timeStampCmd = &cobra.Command{
	Use:   "ts",
	Short: "convert timestamp(ms)",
	Long:  "转换时间戳(ms)",
	Run:   timeStampConvertHandler,
}

func timeStampConvertHandler(cmd *cobra.Command, args []string) {
	formatTime := timer.ConvertTimeStamp(ts)
	log.Printf("时间戳:%d -> %s", ts, formatTime)
}

var ts int64

func init() {
	timeStampCmd.Flags().Int64Var(&ts, "ts", time.Now().Unix(), "input timestamp(ms)")
}
