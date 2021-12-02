package main

import (
	"adventofcode-2021/challenges/utils"
	"log"
)

func main() {
	depths := utils.ReadFileToIntSlice("challenges/1/input")
	numIncreases := getNumberOfIncreases(depths)
	log.Printf("Number of increases %d", numIncreases)
}

func getNumberOfIncreases(depths []int) int {
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
