package others

import (
	"log"
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	type args struct {
		elems []int
		less  func(int, int) bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				elems: []int{11, 2, 44, 3, 66},
				less: func(i int, i2 int) bool {
					return i < i2
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.args.elems)
			BubbleSort1(tt.args.elems, tt.args.less)
			log.Println(tt.args.elems)
		})
	}
}
