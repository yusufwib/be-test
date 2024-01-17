//go:build wireinject
// +build wireinject

package infrastructure

import (
	"context"

	"github.com/yusufwib/be-test/config"
	"github.com/yusufwib/be-test/handler"
	"github.com/yusufwib/be-test/repository"
	"github.com/yusufwib/be-test/service"
	mlog "github.com/yusufwib/be-test/utils/logger"
	"github.com/yusufwib/be-test/utils/mvalidator"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type Dependency struct {
	EmployeeHandler handler.IEmployeeHander
}

func NewDependency(context context.Context,
	logger mlog.Logger,
	validator mvalidator.Validator,
	Config *config.ConfigGroup,
	Database *gorm.DB) *Dependency {
	wire.Build(
		setEmployeeHandler,
		wire.Struct(new(Dependency), "*"),
	)
	return nil
}

var setEmployeeHandler = wire.NewSet(
	repository.NewEmployeeRepository,
	service.NewEmployeeService,
	handler.NewEmployeeHandler,
)
