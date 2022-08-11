package global

import (
	"github.com/go-tour/blog_service/pkg/logger"
	"github.com/go-tour/blog_service/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DataBaseSetting *setting.DataBaseSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSetting
)
