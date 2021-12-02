package main

import "testing"

func Test_getNumberOfIncreases(t *testing.T) {
	type args struct {
		depths []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example Data",
			args: args{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfIncreases(tt.args.depths); got != tt.want {
				t.Errorf("getNumberOfIncreases() = %v, want %v", got, tt.want)
			}
		})
	}
}
