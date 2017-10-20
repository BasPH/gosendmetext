package random

import (
	"testing"
)

func TestInts(t *testing.T) {
	nmin, nmax, maxval := 1, 50, 300
	result := Ints(nmin, nmax, maxval)

	if len(result) < nmin || len(result) > nmax {
		t.Errorf("Incorrect array length. Expected [%d,%d] but got %d", nmin, nmax, len(result))
	}

	for _, n := range result {
		if n > maxval {
			t.Errorf("Given max value %d but found %d", maxval, n)
		}
	}
}

func benchmarkInts(nmax int, b *testing.B) {
	nmin, maxval := 1, 300
	for n := 0; n < b.N; n++ {
		Ints(nmin, nmax, maxval)
	}
}

func BenchmarkInts1(b *testing.B)        { benchmarkInts(1, b) }
func BenchmarkInts10(b *testing.B)       { benchmarkInts(10, b) }
func BenchmarkInts100(b *testing.B)      { benchmarkInts(100, b) }
func BenchmarkInts1000(b *testing.B)     { benchmarkInts(1000, b) }
func BenchmarkInts10000(b *testing.B)    { benchmarkInts(10000, b) }
func BenchmarkInts100000(b *testing.B)   { benchmarkInts(100000, b) }
func BenchmarkInts1000000(b *testing.B)  { benchmarkInts(1000000, b) }
func BenchmarkInts10000000(b *testing.B) { benchmarkInts(10000000, b) }
