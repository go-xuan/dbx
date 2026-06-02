package dbx

import (
	"github.com/go-xuan/configx"
	"github.com/go-xuan/utilx/errorx"
	log "github.com/sirupsen/logrus"
)

func init() {
	RegisterClientBuilder("gorm", GormClientBuilder) // 注册gorm客户端构建器
}

func Initialize() error {
	logger := log.WithField("package", "dbx")
	if err := configx.LoadConfigurator(&Configs{}); err == nil && Initialized() {
		logger.Info("initialize success")
		return nil
	}
	if err := configx.LoadConfigurator(&Config{}); err == nil && Initialized() {
		logger.Info("initialize success")
		return nil
	}
	logger.Warn("initialize failed")
	return errorx.New("failed to initialize dbx")
}
