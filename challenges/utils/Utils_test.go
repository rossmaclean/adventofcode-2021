package utils

import (
	"testing"
)

func TestConvertBinaryToDecimal1(t *testing.T) {
	type args struct {
		binary []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Convert binary to decimal",
			args: args{
				binary: []int{0, 0, 1, 0, 0},
			},
			want: 4,
		},
		{
			name: "Convert binary to decimal",
			args: args{
				binary: []int{1, 1, 1, 1, 0},
			},
			want: 30,
		},
		{
			name: "Convert binary to decimal",
			args: args{
				binary: []int{0, 1, 0, 1, 0},
			},
			want: 10,
		},
		{
			name: "Convert binary to decimal",
			args: args{
				binary: []int{0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0},
			},
			want: 288,
		},
		{
			name: "Convert binary to decimal",
			args: args{
				binary: []int{0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1},
			},
			want: 41,
		},
		{
			name: "Convert binary to decimal",
			args: args{
				binary: []int{1, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1},
			},
			want: 2753,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertBinaryToDecimal(tt.args.binary); got != tt.want {
				t.Errorf("ConvertBinaryToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
