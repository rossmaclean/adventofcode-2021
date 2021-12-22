package utils

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ReadFileToIntSlice(filename string) []int {
	return convertStringSliceToIntSlice(ReadFileToStringSlice(filename))
}

func ReadFileToStringSlice(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\n")
}

func ReadFileForBingo(filename string) ([]int, [][][]int) {
	lines := ReadFileToStringSlice(filename)
	numbers := convertCsvStringToIntSlice(lines[0])
	var boards [][][]int

	for i := 2; i < len(lines)-5; i += 6 {
		var linesForBoard []string
		for y := i; y < i+5; y++ {
			linesForBoard = append(linesForBoard, lines[y])
		}
		boards = append(boards, createBoardFromStringSlice(linesForBoard))
	}
	return numbers, boards
}

func createBoardFromStringSlice(lines []string) [][]int {
	var board [][]int

	for _, line := range lines {
		board = append(board, convertSpaceSeperatedStringToIntSlice(line))
	}
	return board
}

func convertSpaceSeperatedStringToIntSlice(input string) []int {
	reg := regexp.MustCompile("[\\s]{1,2}")
	return convertStringSliceToIntSlice(removeAllMatching(reg.Split(input, -1), ""))
}

func removeAllMatching(slice []string, toMatch string) []string{
	var new []string
	for _, item := range slice {
		if item != toMatch {
			new = append(new, item)
		}
	}
	return new
}

func convertCsvStringToIntSlice(input string) []int {
	return convertStringSliceToIntSlice(strings.Split(input, ","))
}

func convertStringSliceToIntSlice(input []string) []int {
	var convertedValues []int
	for _, s := range input {
		intVal, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		convertedValues = append(convertedValues, intVal)
	}
	return convertedValues
}

func convertStringToIntSlice(input string) []int {
	var convertedValues []int
	for _, char := range input {
		intVal, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatal(err)
		}
		convertedValues = append(convertedValues, intVal)
	}
	return convertedValues
}

func ConvertBinaryToDecimal(binary []int) int {
	var mappings []int
	for i := 0; i < len(binary); i++ {
		if i == 0 {
			mappings = append(mappings, 1)
		} else {
			mappings = append(mappings, mappings[i-1]*2)
		}
	}

	value := 0
	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == 1 {
			value += mappings[(len(binary)-1)-i]
		}
	}
	return value
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
