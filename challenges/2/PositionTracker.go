package main

import (
	"adventofcode-2021/challenges/utils"
	"log"
	"strconv"
	"strings"
)

func main() {
	movements := utils.ReadFileToStringSlice("challenges/2/input")

	result := performMovementsAndGetHorizontalTimesDepth(movements)
	log.Printf("Horizontal times depth %d", result)
	resultWithAim := performMovementsAndGetHorizontalTimesDepthWithAim(movements)
	log.Printf("Horizontal times depth with aim %d", resultWithAim)
}

func performMovementsAndGetHorizontalTimesDepth(movements []string) int {
	var horizontal = 0
	var depth = 0

	for _, movement := range movements {
		step := strings.Split(movement, " ")

		switch step[0] {
		case "forward":
			horizontal += strToInt(step[1])
		case "up":
			depth -= strToInt(step[1])
		case "down":
			depth += strToInt(step[1])
		}
	}
	return horizontal * depth
}

func performMovementsAndGetHorizontalTimesDepthWithAim(movements []string) int {
	var horizontal = 0
	var depth = 0
	var aim = 0

	for _, movement := range movements {
		step := strings.Split(movement, " ")

		switch step[0] {
		case "forward":
			amount := strToInt(step[1])
			horizontal += amount
			depth += aim * amount
		case "up":
			aim -= strToInt(step[1])
		case "down":
			aim += strToInt(step[1])
		}
	}
	return horizontal * depth
}

func strToInt(str string) int {
	amount, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return amount
}
