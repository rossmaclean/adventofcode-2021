package utils

import (
	"io/ioutil"
	"log"
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
