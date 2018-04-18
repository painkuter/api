package config

import (
	"github.com/google/logger"
	"github.com/spf13/viper"
)

// Config contains configuration variables
type Config struct {
	// Debug and release
	IsDebug   bool
	LogDir    string
	LogPrefix string
	LogLevel  int
}

// InitConfig loads config from file
func InitConfig(file string) *Config {

	loadFrom(file)

	return &Config{
		IsDebug:   viper.GetBool("app.isDebug"),
		LogDir:    viper.GetString("app.logDir"),
		LogPrefix: viper.GetString("app.logPrefix"),
		LogLevel:  viper.GetInt("app.logLevel"),
	}
}

func loadFrom(filePath string) {

	viper.SetConfigFile(filePath)
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("Unable to read the file: %s", err)
	}

}
