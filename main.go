package main

import (
	"github.com/Higins/blocks/di"
	"github.com/Higins/blocks/repository"
	"github.com/Higins/blocks/setting"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(logrus.Fields{"type": "main"})

func main() {
	// config
	cfg, err := setting.InitConfig()
	if err != nil {
		log.Error("error parsing config file: ", err)
		return
	}

	// logger
	logLevel, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Warn("error parsing LOG_LEVEL configuration value")
		logLevel = logrus.InfoLevel
	}
	log.Info("logger initialized, log level: ", logLevel)

	gormdb, err := repository.RetryConnectGorm(cfg.DBConnectionString(), 10)
	if err != nil {
		panic(err)
	}
	log.Info("database connection initialized")
	// init router
	applicationRouter := di.InitializeRouter(cfg, gormdb)

	// create error group to detect failed services
	errorGroup := make(chan error, 1)

	// api service
	log.Info("graphql service starting")
	go func() {
		errorGroup <- applicationRouter.InitGraphqlServer(cfg.ProdMode).Run(cfg.ApiHost)
	}()

	// wait until error
	if err := <-errorGroup; err != nil {
		// fatal trigger exit 1
		log.Fatal(err)
	}
}
