package main

import (
	"adventofcode-2021/challenges/utils"
	"log"
)

func main() {
	numbers, boards := utils.ReadFileForBingo("challenges/4/input")

	scoreOfFirstWinningBoard := calculateScoreOfFirstWinningBoard(numbers, boards)
	log.Printf("Score of first winning board %d", scoreOfFirstWinningBoard)

	scoreOfLastWinningBoard := calculateScoreOfLastWinningBoard(numbers, boards)
	log.Printf("Score of last winning board %d", scoreOfLastWinningBoard)
}

func calculateScoreOfFirstWinningBoard(numbers []int, boards [][][]int) int {
	var drawnNumbers []int
	for i := 0; i < 50; i++ {
		for _, board := range boards {
			if hasBoardWon(drawnNumbers, board) {
				return calculateBoardScore(board, drawnNumbers)
			}
		}
		drawnNumbers = append(drawnNumbers, numbers[i])
	}
	return 0
}

func calculateScoreOfLastWinningBoard(numbers []int, boards [][][]int) int {
	var drawnNumbers []int
	for _, number := range numbers {
		if len(boards) > 1 {
			for i, board := range boards {
				if hasBoardWon(drawnNumbers, board) {
					boards = removeBoardFromSliceByIndex(boards, i)
				}
			}
		} else {
			return calculateBoardScore(boards[0], drawnNumbers)
		}
		drawnNumbers = append(drawnNumbers, number)
	}

	return 0
}

func calculateScoreOfLastWinningBoardRecursively(numbers []int, boards[][][]int, drawnNumbers []int) int {
	if len(boards) == 1 {
		return calculateBoardScore(boards[0], drawnNumbers)
	} else {
		drawnNumbers = append(drawnNumbers, numbers[len(drawnNumbers)])
		return calculateScoreOfLastWinningBoardRecursively(numbers, boards, drawnNumbers)
	}
}

func removeBoardFromSliceByIndex(boards [][][]int, index int) [][][]int {
	newBoards := boards[:index]
	for _, board := range boards[index+1:] {
		newBoards = append(newBoards, board)
	}
	return newBoards
}

func calculateBoardScore(board [][]int, drawnNumbers []int) int {
	var sum int
	for _, row := range board {
		for _, num := range row {
			if !utils.Contains(drawnNumbers, num) {
				sum += num
			}
		}
	}
	return sum * drawnNumbers[len(drawnNumbers)-1]
}

func hasBoardWon(drawnNumbers []int, board [][]int) bool {
	//For each row
	rowCount := 0
	for i := 0; i < len(board); i++ {
		rowCount = 0
		for y := 0; y < len(board[i]); y++ {
			if !utils.Contains(drawnNumbers, board[i][y]) {
				break
			} else {
				rowCount++
			}
		}
		if rowCount == len(board) {
			return true
		}
	}

	colCount := 0
	// For each col
	for i := 0; i < len(board); i++ {
		colCount = 0
		for y := 0; y < len(board[i]); y++ {
			if !utils.Contains(drawnNumbers, board[y][i]) {
				break
			} else {
				colCount++
			}
		}
		if colCount == len(board[i]) {
			return true
		}
	}

	return false
}
