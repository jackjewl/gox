package goxSort

import (
	"log"
	"testing"
)

func TestShellSort(t *testing.T) {
	type args struct {
		elems   []int
		compare func(int, int) int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				elems: []int{22, 60, 3, 98, 2, 111, 2, 333, 42, 25},
				compare: func(i int, i2 int) int {
					return i - i2
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.args.elems)
			ShellSort(tt.args.elems, tt.args.compare)
			log.Println(tt.args.elems)
		})
	}
}

//[22 2 11 64 2 34 60 3]
