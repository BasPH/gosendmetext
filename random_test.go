package main

import (
	"testing"
)

func TestRandomInts(t *testing.T) {
	nmin, nmax, maxval := 1, 50, 300
	result := RandomInts(nmin, nmax, maxval)

	if len(result) < nmin || len(result) > nmax {
		t.Errorf("Incorrect array length. Expected [%d,%d] but got %d", nmin, nmax, len(result))
	}

	for _, n := range result {
		if n > maxval {
			t.Errorf("Given max value %d but found %d", maxval, n)
		}
	}
}

func benchmarkRandomInts(nmax int, b *testing.B) {
	nmin, maxval := 1, 300
	for n := 0; n < b.N; n++ {
		RandomInts(nmin, nmax, maxval)
	}
}

func BenchmarkRandomInts1(b *testing.B)        { benchmarkRandomInts(1, b) }
func BenchmarkRandomInts10(b *testing.B)       { benchmarkRandomInts(10, b) }
func BenchmarkRandomInts100(b *testing.B)      { benchmarkRandomInts(100, b) }
func BenchmarkRandomInts1000(b *testing.B)     { benchmarkRandomInts(1000, b) }
func BenchmarkRandomInts10000(b *testing.B)    { benchmarkRandomInts(10000, b) }
func BenchmarkRandomInts100000(b *testing.B)   { benchmarkRandomInts(100000, b) }
func BenchmarkRandomInts1000000(b *testing.B)  { benchmarkRandomInts(1000000, b) }
func BenchmarkRandomInts10000000(b *testing.B) { benchmarkRandomInts(10000000, b) }
