package others

import (
	"log"
	"testing"
)

func TestLcs(t *testing.T) {
	a := []rune("hello")
	b := []rune("hello tom")
	log.Println(Lcs(a, b))
}

func TestLcs1(t *testing.T) {
	a := []rune("hello")
	b := []rune("hello tom")
	log.Println(Lcs1(a, b))
}
