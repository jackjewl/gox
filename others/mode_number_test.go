package others

import (
	"log"
	"testing"
)

func TestModeNumber(t *testing.T) {
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
				elems: []int{1, 2, 2, 2, 2, 2, 3, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ModeNumber(tt.args.elems)
			log.Println(got, got1)
		})
	}
}
