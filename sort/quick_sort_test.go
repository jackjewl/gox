package slice

import (
	"log"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		elems   []int
		low     int
		high    int
		compare func(int, int) int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				elems: []int{11, 2, 34, 5, 11, 95, 22, 9},
				low:   0,
				high:  8,
				compare: func(i int, i2 int) int {
					return i - i2
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.args.elems)
			QuickSort[int](tt.args.elems, tt.args.low, tt.args.high, tt.args.compare)
			log.Println(tt.args.elems)
		})
	}
}
