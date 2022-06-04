package others

import (
	"log"
	"testing"
)

func TestLcs11(t *testing.T) {
	type args struct {
		a []rune
		b []rune
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				a: []rune("hello"),
				b: []rune("hello world"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(Lcs1(tt.args.a, tt.args.b))
		})
	}
}

func TestLcs2(t *testing.T) {
	type args struct {
		a []rune
		b []rune
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				a: []rune("hello"),
				b: []rune("hello world"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(Lcs2(tt.args.a, tt.args.b))
		})
	}
}
