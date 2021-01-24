package util

import (
	"math/rand"
)


func RangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

