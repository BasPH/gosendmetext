package main

import "testing"

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
