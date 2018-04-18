package log

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"api/app/common/config"

	"github.com/google/logger"
)

// InitLogger initializes logger with config settings
func InitLogger(config *config.Config) *logger.Logger {
	// checking directory exists for temporary logging in file
	if _, err := os.Stat(config.LogDir); os.IsNotExist(err) {
		os.MkdirAll(config.LogDir, os.ModePerm)
	}

	y, month, d := time.Now().Date()
	logName := strconv.Itoa(y) + "_" + month.String() + "_" + strconv.Itoa(d)
	filePath := config.LogDir + "/" + config.LogPrefix + "_" + logName

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 0666 | 0660 ?
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}

	l := logger.Init("Logger", true, true, f)
	if l == nil {
		panic("Error logging initializing.")
	}
	logger.Info("Logger initialized")
	return l
}
