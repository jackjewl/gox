package others

import "testing"

func TestSum1(t *testing.T) {
	type args struct {
		elems []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				elems: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum1(tt.args.elems); got != tt.want {
				t.Errorf("Sum1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum2(t *testing.T) {
	type args struct {
		elems []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "test1",
			args: args{
				elems: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum2(tt.args.elems); got != tt.want {
				t.Errorf("Sum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum3(t *testing.T) {
	type args struct {
		elems []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				elems: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum3(tt.args.elems); got != tt.want {
				t.Errorf("Sum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
