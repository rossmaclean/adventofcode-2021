package main

import (
	"adventofcode-2021/challenges/utils"
	"log"
	"strconv"
)

func main() {
	input := utils.ReadFileToStringSlice("challenges/3/input")

	powerConsumption := calculatePowerConsumptionDecimal(input)
	log.Printf("Power consumption %d", powerConsumption)

	lifeSupportRating := calculateLifeSupportRating(input)

	log.Printf("Life support rating %d", lifeSupportRating)
}

func calculatePowerConsumptionDecimal(readings []string) int {
	readingsInt := convertReadingsToInt(readings)

	gammaRate := calculateGammaRateDecimal(readingsInt)
	epsilonRate := calculateEpsilonRateDecimal(readingsInt)

	return gammaRate * epsilonRate
}

func calculateLifeSupportRating(readings []string) int {
	return calculateOxygenGeneratorRating(readings) * calculateCo2ScrubberRating(readings)
}

func calculateCo2ScrubberRating(readings []string) int {
	readingsInt := convertReadingsToInt(readings)

	for i := 0; i < len(readingsInt[0])-1; i++ {
		if len(readingsInt) == 1 {
			break
		}
		leastCommon := flipBit(getMostCommonBitForIndex(readingsInt, i))
		readingsInt = keepStartingWith(readingsInt, leastCommon, i)
	}

	return utils.ConvertBinaryToDecimal(readingsInt[0])
}

func flipBit(bit int) int {
	if bit == 0 {
		return 1
	} else {
		return 0
	}
}

func calculateOxygenGeneratorRating(readings []string) int {
	readingsInt := convertReadingsToInt(readings)

	for i := 0; i < len(readingsInt[0])-1; i++ {
		mostCommon := getMostCommonBitForIndex(readingsInt, i)
		readingsInt = keepStartingWith(readingsInt, mostCommon, i)
	}

	readingsInt = keepStartingWith(readingsInt, 1, 4)

	return utils.ConvertBinaryToDecimal(readingsInt[0])
}

func keepStartingWith(readings [][]int, toKeep int, index int) [][]int {
	var remaining [][]int
	for _, reading := range readings {
		if reading[index] == toKeep {
			remaining = append(remaining, reading)
		}
	}
	return remaining
}

func calculateEpsilonRateDecimal(readings [][]int) int {
	var epsilonBinary []int
	for i := 0; i < len(readings[0]); i++ {
		mostCommon := getMostCommonBitForIndex(readings, i)
		if mostCommon == 1 {
			epsilonBinary = append(epsilonBinary, 0)
		} else {
			epsilonBinary = append(epsilonBinary, 1)
		}

	}
	return utils.ConvertBinaryToDecimal(epsilonBinary)
}

func calculateGammaRateDecimal(readings [][]int) int {
	var gammaBinary []int
	for i := 0; i < len(readings[0]); i++ {
		gammaBinary = append(gammaBinary, getMostCommonBitForIndex(readings, i))
	}
	return utils.ConvertBinaryToDecimal(gammaBinary)
}

func getMostCommonBitForIndex(readings [][]int, index int) int {
	numberOfZeros := 0
	numberOfOnes := 0

	for _, reading := range readings {
		if reading[index] == 0 {
			numberOfZeros++
		} else {
			numberOfOnes++
		}
	}

	if numberOfZeros > numberOfOnes {
		return 0
	} else {
		return 1
	}
}

func convertReadingsToInt(readings []string) [][]int {
	var readingsInt [][]int
	for _, reading := range readings {
		readingsInt = append(readingsInt, convertLineToIntSlice(reading))
	}
	return readingsInt
}

func convertLineToIntSlice(line string) []int {
	var result []int
	for _, char := range line {
		charInt, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, charInt)
	}
	return result
}
