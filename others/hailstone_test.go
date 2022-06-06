package others

import (
	"log"
	"testing"
)

func TestHailstone(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				n: 13,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(Hailstone(tt.args.n))
		})
	}
}
