package app

import (
	"os"
	"os/signal"

	"api/app/common/config"
	"api/app/common/log"

	"github.com/google/logger"
	"api/app/services"
)

const appConfigFilePath = "config/config.dev.toml"

// Application is struct for main application
type Application struct {
	config   *config.Config
	logger   *logger.Logger
	Services interface{}
}

// InitApp creates Appliaction
func InitApp() {

	App := NewApplication(appConfigFilePath)
	defer App.logger.Close()

	service := services.InitService()


	logger.Infof("App started with config: %#v", App.config)

	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)
	for {
		select {
		case sig := <-exit:
			logger.Infof("Got %s signal. Aborting...\n", sig)
			//TODO: wait all tasks done
			logger.Info("App finished.")
			App.logger.Close()
			os.Exit(1)
		}
	}
}

// NewApplication returns pointer to new Application
func NewApplication(cfgPath string) *Application {
	cfg := config.InitConfig(cfgPath)
	return &Application{
		config: cfg,
		logger: newLogger(cfg),
	}
}

func newLogger(config *config.Config) *logger.Logger {
	return log.InitLogger(config)
}
