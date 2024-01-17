package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/yusufwib/be-test/config"
	"github.com/yusufwib/be-test/datasource"
	"github.com/yusufwib/be-test/infrastructure"
	infrastructure_http "github.com/yusufwib/be-test/infrastructure/http"
	"github.com/yusufwib/be-test/utils/mvalidator"

	mlog "github.com/yusufwib/be-test/utils/logger"
)

const (
	appName = "api"
	version = "1.0.0"
)

func main() {
	cfg := config.LoadConfig("config/.env")
	logger := mlog.New("info", "stdout")

	database := datasource.NewPostgreSQLSession(&cfg.PostgreSQLConfig)
	instDB, _ := (*database).DB()
	defer instDB.Close()

	terminalHandler := make(chan os.Signal, 1)
	signal.Notify(terminalHandler, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	application := infrastructure.App{
		TerminalHandler: terminalHandler,
		Cfg:             cfg,
		Logger:          logger,
		Database:        database,
		Validator:       mvalidator.New(),
		Context:         context.Background(),
	}

	infrastructure_http.NewHttpServer(application)
}
