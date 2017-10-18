package main

import (
	"go.uber.org/zap"
	"io/ioutil"
	"bytes"
)

type Words struct {
	data [][]byte
	logger *zap.SugaredLogger
}

func (w *Words) LoadData(textfile string) {
	d, err := ioutil.ReadFile(textfile)
	if err != nil {
		logger.Fatal(err.Error())
	}
	w.data = bytes.Split(d, []byte{'\n'})
}
