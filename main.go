package main

import (
	"go.uber.org/zap"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug = kingpin.Flag("debug", "Debug logging").Short('d').Default("false").Bool()

	address = kingpin.Flag("address", "Address (default 127.0.0.1)").Short('a').Default("127.0.0.1").String()
	port = kingpin.Flag("port", "Port (default 9999)").Short('p').Default("9999").Int()
	textfile = kingpin.Flag("textfile", "Textfile containing words (default /usr/share/dict/words)").Short('t').Default("/usr/share/dict/words").ExistingFile()
	sleep = kingpin.Flag("sleep", "Sleep period (default 0.7s)").Short('s').Default("0.7s").Duration()

	logger *zap.SugaredLogger
)

func init() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	logger = zapLogger.Sugar()
}

func main() {
	kingpin.Parse()

	logger.Infof("Parsed CLI flags",
		zap.Bool("debug", *debug),
		zap.String("address", *address),
		zap.Int("port", *port),
		zap.String("textfile", *textfile),
		zap.Duration("sleep", *sleep),
	)
}