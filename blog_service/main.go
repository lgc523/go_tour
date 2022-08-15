package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	validatorV10 "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
	"github.com/go-tour/blog_service/global"
	"github.com/go-tour/blog_service/internal/model"
	"github.com/go-tour/blog_service/internal/routers"
	"github.com/go-tour/blog_service/pkg/logger"
	"github.com/go-tour/blog_service/pkg/setting"
	"github.com/go-tour/blog_service/pkg/tracer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	err = setUpZapLog()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	zap.L().Info("zapLog setup...")
	global.Log.Info("global zapLog setup...")
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupValidator()
	if err != nil {
		log.Fatalf("init.setupValidator err: %v", err)
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

func setUpZapLog() error {
	appSetting := global.AppSetting
	writeSyncer := getLogWriter(appSetting.LogSavePath+"/"+appSetting.LogFileName+appSetting.LogFileExt, appSetting.MaxLogFileSize, appSetting.MaxLogBackUp)

	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(appSetting.LogLevel))
	if err != nil {
		return err
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg := zap.New(core, zap.AddCaller())
	//zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	global.Log = lg
	return nil
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
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

func setupValidator() error {
	//global.Validator, = validator.NewCustomValidator()
	//global.Validator.Engine()
	//binding.Validator = global.Validator
	uni := ut.New(zh.New(), en.New(), zh_Hant_TW.New(), ja.New())
	v, ok := binding.Validator.Engine().(*validatorV10.Validate)
	if ok {
		zhTran, _ := uni.GetTranslator("zh")
		enTran, _ := uni.GetTranslator("en")
		tw, _ := uni.GetTranslator("zh_Hant_tw")
		ja, _ := uni.GetTranslator("ja")
		err := zh_translations.RegisterDefaultTranslations(v, zhTran)
		if err != nil {
			return err
		}
		err = tw_translations.RegisterDefaultTranslations(v, tw)
		if err != nil {
			return err
		}
		err = en_translations.RegisterDefaultTranslations(v, enTran)
		if err != nil {
			return err
		}
		err = ja_translations.RegisterDefaultTranslations(v, ja)
		if err != nil {
			return err
		}
	}
	global.Ut = uni
	return nil
}
