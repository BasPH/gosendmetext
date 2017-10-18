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
	w.logger.Debugf("Reading file %s", textfile)

	d, err := ioutil.ReadFile(textfile)
	if err != nil {
		logger.Fatal(err.Error())
	}
	split := bytes.Split(d, []byte{'\n'})

	var result [][]byte
	for _, s := range split {
		if len(s) > 0 {
			result = append(result, s)
		}
	}
	w.data = result

	w.logger.Debugf("Read file %s with %v lines", textfile, len(w.data))
}
