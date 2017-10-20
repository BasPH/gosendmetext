package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/BasPH/gosendmetext/words"
	"net"
	"os"
	"time"
)

var (
	debug = kingpin.Flag("debug", "Debug logging").Short('d').Default("false").Bool()

	address  = kingpin.Flag("address", "Address (default 127.0.0.1)").Short('a').Default("127.0.0.1").String()
	port     = kingpin.Flag("port", "Port (default 9999)").Short('p').Default("9999").Int()
	protocol = kingpin.Flag("protocol", "Protocol (default TCP)").Short('n').Default("tcp").Enum("tcp", "udp")
	textfile = kingpin.Flag("textfile", "Textfile containing words, one per line (default /usr/share/dict/words)").Short('t').Default("/usr/share/dict/words").ExistingFile()
	minwords = kingpin.Flag("minwords", "Minimum number of words to send each time (default 1)").Default("1").Int()
	maxwords = kingpin.Flag("maxwords", "Maximum number of words to send each time (default 50)").Default("50").Int()
	sleep    = kingpin.Flag("sleep", "Sleep period (default 0.7s)").Short('s').Default("0.7s").Duration()

	log *logrus.Logger
)

func init() {
	log = logrus.New()
	log.Formatter = logrus.Formatter(&logrus.JSONFormatter{})
	log.Out = os.Stdout
	log.Level = logrus.InfoLevel
}

func main() {
	kingpin.Parse()
	if *debug {
		log.Level = logrus.DebugLevel
	}

	log.WithFields(logrus.Fields{
		"debug":    *debug,
		"address":  *address,
		"port":     *port,
		"protocol": *protocol,
		"textfile": *textfile,
		"minwords": *minwords,
		"maxwords": *maxwords,
		"sleep":    *sleep,
	}).Info("Parsed CLI flags")

	w := words.NewWords(log, *textfile)

	conn, err := net.Dial(*protocol, fmt.Sprintf("%s:%d", *address, *port))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	for {
		w.SendRandomWords(conn, *minwords, *maxwords)
		time.Sleep(*sleep)
	}
}
