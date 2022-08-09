package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/internal/model"
	"github.com/go-tour/blog_service/internal/routers"
	"github.com/go-tour/blog_service/pkg/logger"
	"github.com/go-tour/blog_service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setUpLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func main() {
	global.Logger.InfoF("[bolg_service server starting...]")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	server := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeOut * time.Second,
		MaxHeaderBytes: 2 << 20,
	}
	global.Logger.InfoF("[bolg_service server start over, listened on port %s...]", global.ServerSetting.HttpPort)
	err := server.ListenAndServe()
	global.Logger.PanicF("engine start error: %s", err.Error())
}

func setupSetting() error {
	vpSet, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = vpSet.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = vpSet.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = vpSet.ReadSection("Database", &global.DataBaseSetting)
	if err != nil {
		return err
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
			Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
			MaxSize:   523,
			MaxAge:    15,
			LocalTime: true,
		},
		global.AppSetting.LogPrefix,
		log.LstdFlags).WithCaller(2)
	return nil
}
