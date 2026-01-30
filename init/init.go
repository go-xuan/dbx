package init

import (
	"github.com/go-xuan/configx"
	"github.com/go-xuan/dbx"
	"github.com/go-xuan/utilx/errorx"
)

func init() {
	errorx.Panic(Init())
}

func Init() error {
	var err error
	if err = configx.LoadConfigurator(&dbx.Configs{}); err == nil && dbx.Initialized() {
		return nil
	} else if err = configx.LoadConfigurator(&dbx.Config{}); err == nil && dbx.Initialized() {
		return nil
	}
	return errorx.Wrap(err, "init database failed")
}
