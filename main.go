package main

import (
	"go.uber.org/zap"
	"gopkg.in/alecthomas/kingpin.v2"
	"net"
	"fmt"
)

var (
	debug = kingpin.Flag("debug", "Debug logging").Short('d').Default("false").Bool()

	address = kingpin.Flag("address", "Address (default 127.0.0.1)").Short('a').Default("127.0.0.1").String()
	port = kingpin.Flag("port", "Port (default 9999)").Short('p').Default("9999").Int()
	protocol = kingpin.Flag("protocol", "Protocol (default TCP)").Short('n').Default("tcp").Enum("tcp", "udp")
	textfile = kingpin.Flag("textfile", "Textfile containing words, one per line (default /usr/share/dict/words)").Short('t').Default("/usr/share/dict/words").ExistingFile()
	minwords = kingpin.Flag("minwords", "Minimum number of words to send each time (default 1)").Default("1").Int()
	maxwords = kingpin.Flag("maxwords", "Maximum number of words to send each time (default 50)").Default("50").Int()
	sleep = kingpin.Flag("sleep", "Sleep period (default 0.7s)").Short('s').Default("0.7s").Duration()

	logger *zap.SugaredLogger
)

func main() {
	kingpin.Parse()

	initLogging()
	logger.Infof("Parsed CLI flags",
		zap.Bool("debug", *debug),
		zap.String("address", *address),
		zap.Int("port", *port),
		zap.String("protocol", *protocol),
		zap.String("textfile", *textfile),
		zap.Int("minwords", *minwords),
		zap.Int("maxwords", *maxwords),
		zap.Duration("sleep", *sleep),
	)

	w := NewWords(logger, *textfile)

	conn, err := connect(*protocol, *address, *port)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer conn.Close()

	result, _ := w.RandomWords(*minwords, *maxwords)
	conn.Write(result)
}

func initLogging() {
	cfg := zap.NewProductionConfig()
	if *debug {
		cfg.Level.SetLevel(zap.DebugLevel)
	}
	cfg.Build()
	zapLogger, _ := cfg.Build()
	defer zapLogger.Sync()
	logger = zapLogger.Sugar()
}

func connect(protocol string, host string, port int) (net.Conn, error) {
	return net.Dial(protocol, fmt.Sprintf("%s:%d", host, port))
}