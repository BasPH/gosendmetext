package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInts(nmin int, nmax int, maxval int) []int {
	nsize := 1
	if nmax > 1 {
		nsize = rand.Intn(nmax-nmin) + nmin
	}
	result := make([]int, nsize)
	for i := 0; i <= nsize-1; i++ {
		result[i] = rand.Intn(maxval)
	}
	return result
}
