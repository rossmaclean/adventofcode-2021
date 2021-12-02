package main

import (
	"adventofcode-2021/challenges/utils"
	"log"
)

func main() {
	depths := utils.ReadFileToIntSlice("challenges/1/input")

	numIncreasesIndividual := getNumberOfIncreasesIndividual(depths)
	log.Printf("Number of increases individual %d", numIncreasesIndividual)

	numToSum := 3
	numIncreasesSum := getNumberOfIncreasesSum(depths, numToSum)
	log.Printf("Number of increases sum using %d values %d", numToSum, numIncreasesSum)
}

func getNumberOfIncreasesIndividual(depths []int) int {
	var increases = 0
	for i := range depths {
		if i == 0 {
			continue
		}
		if depths[i] > depths[i-1] {
			increases++
		}
	}
	return increases
}

func getNumberOfIncreasesSum(depths []int, numToSum int) int {
	var increases = 0
	for i := 0; i < len(depths)-numToSum; i++ {
		firstSum := getSum(depths, i, numToSum)
		secondsSum := getSum(depths, i+1, numToSum)

		if secondsSum > firstSum {
			increases++
		}
	}
	return increases
}

func getSum(depths []int, startIndex int, numToSum int) int {
	var sum = 0
	for i := 0; i < numToSum; i++ {
		sum += depths[startIndex+i]
	}
	return sum
}
