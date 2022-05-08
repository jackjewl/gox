package vector

import (
	"log"
	"math/rand"
	"testing"
)

func TestVector_Permute(t *testing.T) {

	for i := 0; i < 10; i++ {
		log.Println(i, rand.Int())
	}
}
