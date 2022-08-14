package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/internal/model"
	"github.com/go-tour/blog_service/internal/routers"
	"github.com/go-tour/blog_service/pkg/logger"
	"github.com/go-tour/blog_service/pkg/setting"
	"github.com/go-tour/blog_service/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	port    string
	runMode string
	config  string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init flag err:%v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setUpLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}

}

func setupFlag() error {
	flag.StringVar(&port, "port", "5023", "启动端口")
	flag.StringVar(&runMode, "mode", "release", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定撇脂文件路径")
	flag.Parse()
	return nil
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	server := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeOut * time.Second,
		MaxHeaderBytes: 2 << 20,
	}
	err := server.ListenAndServe()
	global.Logger.PanicF("engine start error: %s", err.Error())
}

func setupSetting() error {
	vpSet, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = vpSet.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = vpSet.ReadSection("App", &global.AppSetting)
	global.AppSetting.DefaultContextTimeout *= time.Second
	if err != nil {
		return err
	}

	err = vpSet.ReadSection("Database", &global.DataBaseSetting)
	if err != nil {
		return err
	}

	err = vpSet.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = vpSet.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = vpSet.ReadSection("DingTalk", &global.DingTalkSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setUpLogger() error {
	global.Logger = logger.NewLogger(
		&lumberjack.Logger{
			Filename:   global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
			MaxSize:    10,
			MaxBackups: 2,
			MaxAge:     15,
			LocalTime:  true,
			Compress:   false,
		},
		global.AppSetting.LogPrefix,
		log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(global.AppSetting.AppName,
		"c1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
