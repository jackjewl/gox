package apply

import (
	statck "gox/stack"
	"log"
	"testing"
)

func TestBaseConvert(t *testing.T) {
	s := statck.NewLinkedStack[rune]()
	BaseConvert(s, 100, 16)

	for !s.Empty() {
		log.Printf("%c\t", s.Pop())
	}
}

func TestParenthesisMatch(t *testing.T) {
	expression := "(2+3)*5+}[5-2]"
	r := ParenthesisMatch([]rune(expression))
	log.Println(r)
}
