package bitmap

import (
	"log"
	"testing"
)

func TestBitmap_Exist(t *testing.T) {
	b := NewBitmap(10)
	b.Set(1)
	b.Set(2)
	log.Println(b.Exist(0))
	log.Println(b.Exist(1))
	log.Println(b.Exist(2))
	log.Println(b.Exist(10))
}
