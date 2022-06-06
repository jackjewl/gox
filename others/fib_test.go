package others

import (
	"log"
	"reflect"
	"testing"
)

func TestFib1(t *testing.T) {
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
				n: 40,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fib1(tt.args.n)
			log.Println(got)
		})
	}
}

func TestFib2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				n: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(Fib2(tt.args.n))
		})
	}
}

func TestFib3(t *testing.T) {
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
				n: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(Fib3(tt.args.n))
		})
	}
}

func TestFib3Part(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *SecondOrderMatrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib3Part(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fib3Part() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecondOrderMatrixMutl(t *testing.T) {
	type args struct {
		s1 *SecondOrderMatrix
		s2 *SecondOrderMatrix
	}
	tests := []struct {
		name string
		args args
		want *SecondOrderMatrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecondOrderMatrixMutl(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SecondOrderMatrixMutl() = %v, want %v", got, tt.want)
			}
		})
	}
}
