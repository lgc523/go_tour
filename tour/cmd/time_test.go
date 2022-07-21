package cmd

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Logf("时间戳:%d", time.Now().Unix())
	t.Logf("纳秒时间戳:%d", time.Now().UnixNano())
	t.Logf("毫秒时间戳:%d", time.Now().UnixMicro())
	t.Logf("毫秒时间戳:%d", time.Now().UnixMilli())
	t.Logf("毫秒时间戳:%d", time.Now().UnixNano()/1e6)
	t.Logf("格式化:%s", time.Now().Format("2006-01-02 15-04-05"))
	t.Logf("格式化时间戳:%s", time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15-04-05"))
	t.Log(time.Date(2022, 7, 22, 2, 30, 30, 30, time.Local))
}
