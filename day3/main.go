package main

import (
	"aoc/fileloading"
	"fmt"
	"log"
	"strconv"
)

func main() {
		log.SetPrefix("day2: ")
    log.SetFlags(0)

		binaryNumbers, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }

		part1(binaryNumbers)

		fmt.Println(part2(binaryNumbers))
}

func part1(binaryNumbers []string) {
		var totalNumbers = len(binaryNumbers)
		var counters = make([]int, len(binaryNumbers[0]))
		fmt.Println(counters)

		for _, line := range binaryNumbers {
			singleNumber := []rune(line)

			for index, character := range singleNumber {
				counters[index] += int(character-'0') // Converts to int32 instead of the int of the ASCII
			}
		}

		var gamma, epsilon string

		for _, value := range counters {
			if (value > totalNumbers/2) {
				gamma += "1"
				epsilon += "0"
			} else {
				gamma += "0"
				epsilon += "1"
			}
		}


		fmt.Println(gamma)
		fmt.Println(epsilon)

		var gamma_int, _ = strconv.ParseInt(gamma, 2, 32)
		var epsilon_int, _ = strconv.ParseInt(epsilon, 2, 32)

		fmt.Println(gamma_int)
		fmt.Println(epsilon_int)
		fmt.Println(gamma_int*epsilon_int)
}

func part2(data []string) int {
	oxygenString := lifeSupportRating(data, 0, true)
	co2String := lifeSupportRating(data, 0, false)

	oxygen, _ := strconv.ParseInt(oxygenString, 2, 32)
	co2, _ := strconv.ParseInt(co2String, 2, 32)
	return int(oxygen) * int(co2)
}

// Recursive function to find the oxygen or CO2 rating
func lifeSupportRating(data []string, bitIndex int, countMost bool) string {
	if len(data) == 1 {
		return data[0]
	}

	countOfOnes := 0
	var bitFilter byte = '0'
	for _, dataLine := range data {
		if dataLine[bitIndex] == '1' {
			countOfOnes++
		}
		// I can't understand why I need to use len-1, but I have to
		if countOfOnes > (len(data)-1)/2 {
			bitFilter = '1'
			break
		}
	}

	if !countMost {
		if bitFilter == '1' {
			bitFilter = '0'
		} else {
			bitFilter = '1'
		}
	}

	var filtered []string
	for _, dataLine := range data {
		if dataLine[bitIndex] == bitFilter {
			filtered = append(filtered, dataLine)
		}
	}

	// Recursion baby!
	return lifeSupportRating(filtered, bitIndex+1, countMost)
}
