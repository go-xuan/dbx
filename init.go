package dbx

import (
	"github.com/go-xuan/configx"
	log "github.com/sirupsen/logrus"
)

func init() {
	RegisterClientBuilder("gorm", GormClientBuilder) // 注册gorm客户端构建器
	Init()                                           // 初始化数据库
}

func Init() {
	logger := log.WithField("package", "dbx")
	if err := configx.LoadConfigurator(&Configs{}); err == nil && Initialized() {
		logger.Info("initialized success")
		return
	}
	if err := configx.LoadConfigurator(&Config{}); err == nil && Initialized() {
		logger.Info("initialized success")
		return
	}
	logger.Warn("initialized failed")
}
