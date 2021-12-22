package main

import (
	"adventofcode-2021/challenges/utils"
	"log"
)

func main() {
	lines := utils.ReadFileToStringSlice("challenges/regex_test/regex_test_input")

	var splitLines [][]string
	for _, line := range lines {
		log.Println(line)
		//splitLines = append(splitLines, splitLineByRegex(line))
	}

	log.Println("=========================")
	for _, line := range splitLines {
		for _, char := range line {
			log.Println(char)
		}
		log.Println("=============")
	}
}

//func splitLineByRegex(line string) []string {
//	exp := regexp.MustCompile("[\\s]{1,2}")
//	return removeAllMatching(exp.Split(line, -1), "")
//}

