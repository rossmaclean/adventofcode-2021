package main

import (
	"adventofcode-2021/challenges/utils"
	"reflect"
	"testing"
)

func Test_performMovementsAndGetHorizontalTimesDepth(t *testing.T) {
	type args struct {
		movements []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test perform movements and get horizontal times depth",
			args: args{
				movements: []string{
					"forward 5",
					"down 5",
					"forward 8",
					"up 3",
					"down 8",
					"forward 2",
				},
			},
			want: 150,
		}, {
			name: "Test perform movements and get horizontal times depth actual input",
			args: args{
				movements: utils.ReadFileToStringSlice("input"),
			},
			want: 1868935,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := performMovementsAndGetHorizontalTimesDepth(tt.args.movements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("performMovementsAndGetHorizontalTimesDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_performMovementsAndGetHorizontalTimesDepthWithAim(t *testing.T) {
	type args struct {
		movements []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test perform movements and get horizontal times depth with aim",
			args: args{
				movements: []string{
					"forward 5",
					"down 5",
					"forward 8",
					"up 3",
					"down 8",
					"forward 2",
				},
			},
			want: 900,
		},
		{
			name: "Test perform movements and get horizontal times depth with aim actual input",
			args: args{
				movements: utils.ReadFileToStringSlice("input"),
			},
			want: 1965970888,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := performMovementsAndGetHorizontalTimesDepthWithAim(tt.args.movements); got != tt.want {
				t.Errorf("performMovementsAndGetHorizontalTimesDepthWithAim() = %v, want %v", got, tt.want)
			}
		})
	}
}
