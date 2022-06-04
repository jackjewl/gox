package others

import (
	"log"
	"testing"
)

func TestReverse1(t *testing.T) {
	type args struct {
		elems []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				elems: []int{1, 2, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.args.elems)
			Reverse1[int](tt.args.elems)
			log.Println(tt.args.elems)
		})
	}
}
