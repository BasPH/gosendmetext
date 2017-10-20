package words

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"sync"
	"net"
	"github.com/BasPH/gosendmetext/random"
)

// Words is a service sending randomly generated sentences
type Words struct {
	data      [][]byte
	log       *logrus.Logger
	ch        chan bool
	waitGroup *sync.WaitGroup
}

// NewWords initializes a new Words instance
func NewWords(log *logrus.Logger, textfile string) *Words {
	w := &Words{
		log: log,
	}
	w.loadData(textfile)
	return w
}

// LoadData reads a file containing words and sets the data field in a Words struct
func (w *Words) loadData(textfile string) {
	w.log.Debugf("Reading file %s", textfile)

	d, err := ioutil.ReadFile(textfile)
	if err != nil {
		w.log.Fatal(err.Error())
	}
	split := bytes.Split(d, []byte{'\n'})

	var result [][]byte
	for _, s := range split {
		if len(s) > 0 {
			result = append(result, s)
		}
	}
	w.data = result

	w.log.Debugf("Read file %s with %v lines", textfile, len(w.data))
}

func (w *Words) randomWords(nmin int, nmax int) ([]byte, error) {
	if nmin > len(w.data) {
		return nil, fmt.Errorf("nmin (%d) must be smaller than data length (%d)", nmin, len(w.data))
	}
	if nmax > len(w.data) {
		return nil, fmt.Errorf("nmax (%d) must be smaller than data length (%d)", nmax, len(w.data))
	}
	if nmin > nmax {
		return nil, fmt.Errorf("nmin must be <= nmax")
	}

	idxs := random.Ints(nmin, nmax, len(w.data))
	var result []byte
	for i := 0; i <= len(idxs)-1; i++ {
		result = append(result, w.data[idxs[i]]...)
		if i+1 != len(idxs) {
			result = append(result, ' ')
		} else {
			result = append(result, '\n')
		}
	}

	w.log.Debugf("Returning %d random words", len(idxs))
	return result, nil
}

// SendRandomWords sends [nmin,nmax] random words to the given connection
func (w *Words) SendRandomWords(c net.Conn, nmin int, nmax int) (n int, err error) {
	rw, _ := w.randomWords(nmin, nmax)
	w.log.Infof("Writing random words: %v", string(rw[:]))
	return c.Write(rw)
}
