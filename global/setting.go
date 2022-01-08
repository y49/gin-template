package global

import (
	"gin-template/pkg/logger"
	"gin-template/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	EmailSetting    *setting.EmailSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JaegerSetting   *setting.JaegerS
)
