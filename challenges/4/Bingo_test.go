package main

import "testing"

func Test_calculateScoreOfFirstWinningBoard(t *testing.T) {
	type args struct {
		drawnNumbers []int
		boards       [][][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Calculate score of first winning board",
			args: args{
				drawnNumbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				boards: [][][]int{
					{
						{22, 13, 17, 11, 0},
						{8, 2, 23, 4, 24},
						{21, 9, 14, 16, 7},
						{6, 10, 3, 18, 5},
						{1, 12, 20, 15, 19},
					},
					{
						{3, 15, 0, 2, 22},
						{9, 18, 13, 17, 5},
						{19, 8, 7, 25, 23},
						{20, 11, 10, 24, 4},
						{14, 21, 16, 12, 6},
					},
					{
						{14, 21, 17, 24, 4},
						{10, 16, 15, 9, 19},
						{18, 8, 23, 26, 20},
						{22, 11, 13, 6, 5},
						{2, 0, 12, 3, 7},
					},
				},
			},
			want: 4512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateScoreOfFirstWinningBoard(tt.args.drawnNumbers, tt.args.boards); got != tt.want {
				t.Errorf("calculateScoreOfFirstWinningBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasBoardWon(t *testing.T) {
	type args struct {
		drawnNumbers []int
		board        [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Has board won 1",
			args: args{
				drawnNumbers: []int{1, 2, 3, 4, 5, 6, 7},
				board: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
			},
			want: false,
		},
		{
			name: "Has board won 2",
			args: args{
				drawnNumbers: []int{1, 2, 3, 4, 5, 6, 7},
				board: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{1, 2, 3, 4, 5},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
			},
			want: true,
		},
		{
			name: "Has board won 3",
			args: args{
				drawnNumbers: []int{1, 2, 3, 4, 5, 6, 7},
				board: [][]int{
					{22, 1, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{1, 3, 9, 4, 5},
					{6, 4, 9, 18, 5},
					{1, 5, 20, 15, 19},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasBoardWon(tt.args.drawnNumbers, tt.args.board); got != tt.want {
				t.Errorf("hasBoardWon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateBoardScore(t *testing.T) {
	type args struct {
		board        [][]int
		drawnNumbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Calculate board score 1",
			args: args{
				board: [][]int{
					{14, 21, 17, 24  ,4},
					{10, 16, 15,  9 ,19},
					{18 , 8, 23, 26 ,20},
					{22, 11, 13,  6  ,5},
					{2 , 0, 12,  3  ,7},
				},
				drawnNumbers: []int{7,4,9,5,11,17,23,2,0,14,21,24},
			},
			want: 4512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateBoardScore(tt.args.board, tt.args.drawnNumbers); got != tt.want {
				t.Errorf("calculateBoardScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateScoreOfLastWinningBoard(t *testing.T) {
	type args struct {
		numbers []int
		boards  [][][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Calculate score of last winning board",
			args: args{
				numbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				boards: [][][]int{
					{
						{22, 13, 17, 11, 0},
						{8, 2, 23, 4, 24},
						{21, 9, 14, 16, 7},
						{6, 10, 3, 18, 5},
						{1, 12, 20, 15, 19},
					},
					{
						{3, 15, 0, 2, 22},
						{9, 18, 13, 17, 5},
						{19, 8, 7, 25, 23},
						{20, 11, 10, 24, 4},
						{14, 21, 16, 12, 6},
					},
					{
						{14, 21, 17, 24, 4},
						{10, 16, 15, 9, 19},
						{18, 8, 23, 26, 20},
						{22, 11, 13, 6, 5},
						{2, 0, 12, 3, 7},
					},
				},
			},
			want: 1924,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateScoreOfLastWinningBoard(tt.args.numbers, tt.args.boards); got != tt.want {
				t.Errorf("calculateScoreOfLastWinningBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateScoreOfLastWinningBoardRecursively(t *testing.T) {
	type args struct {
		numbers      []int
		boards       [][][]int
		drawnNumbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Calculate score of last winning board",
			args: args{
				numbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				boards: [][][]int{
					{
						{22, 13, 17, 11, 0},
						{8, 2, 23, 4, 24},
						{21, 9, 14, 16, 7},
						{6, 10, 3, 18, 5},
						{1, 12, 20, 15, 19},
					},
					{
						{3, 15, 0, 2, 22},
						{9, 18, 13, 17, 5},
						{19, 8, 7, 25, 23},
						{20, 11, 10, 24, 4},
						{14, 21, 16, 12, 6},
					},
					{
						{14, 21, 17, 24, 4},
						{10, 16, 15, 9, 19},
						{18, 8, 23, 26, 20},
						{22, 11, 13, 6, 5},
						{2, 0, 12, 3, 7},
					},
				},
			},
			want: 1924,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateScoreOfLastWinningBoardRecursively(tt.args.numbers, tt.args.boards, tt.args.drawnNumbers); got != tt.want {
				t.Errorf("calculateScoreOfLastWinningBoardRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}