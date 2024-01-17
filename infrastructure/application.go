package infrastructure

import (
	"context"
	"os"

	"github.com/yusufwib/be-test/config"
	"github.com/yusufwib/be-test/utils/mvalidator"

	"gorm.io/gorm"

	mlog "github.com/yusufwib/be-test/utils/logger"
)

type App struct {
	TerminalHandler chan os.Signal
	Cfg             *config.ConfigGroup
	Logger          mlog.Logger
	Database        *gorm.DB
	Validator       mvalidator.Validator
	Context         context.Context
}
