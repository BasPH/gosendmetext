package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"net"
	"os"
	"bufio"
)

var (
	debug = kingpin.Flag("debug", "Debug logging").Short('d').Default("false").Bool()

	address  = kingpin.Flag("address", "Address (default 127.0.0.1)").Short('a').Default("127.0.0.1").String()
	port     = kingpin.Flag("port", "Port (default 9999)").Short('p').Default("9999").Int()
	protocol = kingpin.Flag("protocol", "Protocol (default TCP)").Short('n').Default("tcp").Enum("tcp", "udp")
	sleep    = kingpin.Flag("sleep", "Sleep period (default 0s)").Short('s').Default("0s").Duration()

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
		"sleep":    *sleep,
	}).Info("Parsed CLI flags")

	l, _ := net.Listen(*protocol, fmt.Sprintf("%s:%d", *address, *port))
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		log.Errorf("error accepting connection: %v", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	log.Infof("Listening on %s %s:%d", *protocol, *address, *port)
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Debugf("Message received: %s", message[:len(message)-1])
	}
}
