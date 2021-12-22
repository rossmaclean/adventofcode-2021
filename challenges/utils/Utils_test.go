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

func TestContains(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Contains 1",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				e: 8,
			},
			want: false,
		},
		{
			name: "Contains 2",
			args: args{
				s: []int{5, 7, 8, 2, 1},
				e: 3,
			},
			want: false,
		},
		{
			name: "Contains 3",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				e: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
