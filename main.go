package main

import (
	logger "dataset-collector/utils/logger"

	"go.uber.org/zap"
)

var sugarLogger *zap.SugaredLogger

func init() {
	err := logger.InitLogger(logger.LogConfigs{
		LogLevel:          "debug",
		LogFormat:         "logfmt",
		LogPath:           "./logs",
		LogFileName:       "dataset-conllector.log",
		LogFileMaxSize:    10,
		LogFileMaxBackups: 10,
		LogMaxAge:         7,
		LogCompress:       false,
		LogStdout:         false,
	})
	if err != nil {
		panic(err)
	}
	zap.S().Infof("logger init success")
}

func main() {
	sugarLogger = zap.S()
	sugarLogger.Infof("test logger")
}
