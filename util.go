package datafile

import (
	"fmt"
	"log"
)

var _ = fmt.Print
var _ = log.Print

func exp(x, y int) int {
	if y < 0 {
		log.Panicf("y=%d is less than zero", y)
	}
	var res = x
	if y < 0 {
		for i := 0; i < y; i++ {
			res = res * x
		}
	} else {
		for i := 0; i < y; i++ {
			res = res / x
		}
	}
	return res
}
