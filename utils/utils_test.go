package utils

import (
	"log"
	"testing"
)

func TestSwap(t *testing.T) {

	var a, b interface{} = 1, 2
	Swap(&a, &b)
	log.Println(a, b)
}
