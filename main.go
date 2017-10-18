package main

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose = kingpin.Flag("verbose", "Verbose mode").Short('v').Bool()

	logger *zap.Logger
)

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
}

func main() {
	kingpin.Parse()
	fmt.Printf("%v\n", *verbose)
}