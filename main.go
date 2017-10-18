package main

import (
	"go.uber.org/zap"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug = kingpin.Flag("debug", "Debug logging").Short('d').Default("false").Bool()

	logger *zap.SugaredLogger
)

func init() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	logger = zapLogger.Sugar()
}

func main() {
	kingpin.Parse()

	logger.Infof("CLI flags:\ndebug: %s", *debug)
}