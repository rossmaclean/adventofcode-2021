package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadFileToIntSlice(filename string) []int {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return convertStringSliceToIntSlice(strings.Split(string(b), "\n"))
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