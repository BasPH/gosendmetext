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

func NewWords(logger *zap.SugaredLogger, textfile string) *Words {
	w := &Words{
		logger: logger,
	}
	w.LoadData(textfile)
	return w
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

func (w *Words) RandomWords(nmin int, nmax int) []byte {

	idxs := RandomInts(nmin, nmax, len(w.data))
	var result []byte
	for i := 0; i <= len(idxs)-1; i++ {
		result = append(result, w.data[idxs[i]]...)
		if i+1 != len(idxs) {
			result = append(result, ' ')
		} else {
			result = append(result, '\n')
		}
	}

	w.logger.Debugf("Returning %d random words", len(idxs))
	return result
}
