package main

import "testing"

func Test_getNumberOfIncreasesIndividual(t *testing.T) {
	type args struct {
		depths []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Get number of increases individual",
			args: args{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfIncreasesIndividual(tt.args.depths); got != tt.want {
				t.Errorf("getNumberOfIncreasesIndividual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNumberOfIncreasesSum(t *testing.T) {
	type args struct {
		depths   []int
		numToSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Get number of increases sum",
			args: args{
				depths:   []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
				numToSum: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfIncreasesSum(tt.args.depths, tt.args.numToSum); got != tt.want {
				t.Errorf("getNumberOfIncreasesSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
