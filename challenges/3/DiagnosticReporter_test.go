package main

import (
	"adventofcode-2021/challenges/utils"
	"testing"
)

func Test_calculatePowerConsumptionDecimal(t *testing.T) {
	type args struct {
		readings []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Calculate power consumption decimal",
			args: args{
				readings: []string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			want: 198,
		},
		{
			name: "Calculate power consumption decimal actual input",
			args: args{
				readings: utils.ReadFileToStringSlice("input"),
			},
			want: 3969000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatePowerConsumptionDecimal(tt.args.readings); got != tt.want {
				t.Errorf("calculatePowerConsumptionDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateLifeSupportRating(t *testing.T) {
	type args struct {
		readings []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Calculate life support rating",
			args: args{
				readings: []string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			want: 230,
		},
		{
			name: "Calculate life support rating actual input",
			args: args{
				readings: utils.ReadFileToStringSlice("input"),
			},
			want: 4267809,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateLifeSupportRating(tt.args.readings); got != tt.want {
				t.Errorf("calculateLifeSupportRating() = %v, want %v", got, tt.want)
			}
		})
	}
}
