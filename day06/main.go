package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
		log.SetPrefix("day6: ")
    log.SetFlags(0)

		fish, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }

		fmt.Println("Day 6")
		utils.PrintAnswer(1, part1(fish))
		utils.PrintAnswer(2, part2(fish))
}

func part1(data []string) int {
	var DAYS = 80
	var fish = getFish(data)

	for i := 1; i <= DAYS; i++ {
		for fi, v := range fish {
			// new fish spawned
			if fish[fi] == 0 {
				fish = append(fish, 8)
			}

			fish[fi] = ageFish(v)
		}

		// fmt.Printf("Day %d: ", i)
		// fmt.Print(fish)
		// fmt.Println()

	}
	return len(fish)
}

func part2(data []string) int {
	var DAYS = 256
	var fish = getFish(data)

	var ages = make([]int, 9)
	// change to age array
	for _, v := range fish {
		ages[v]++
	}

	for i := 1; i <= DAYS; i++ {
		var newAges = make([]int, 9)
		for age := 8; age >= 0; age-- {
			if age == 0 {
				newAges[6] += ages[0]
				newAges[8] += ages[0]
				continue
			}

			newAges[age - 1] = ages[age]
		}

		ages = newAges
	}

	sumFish := 0
	for _, num := range ages {
			sumFish += num
	}
	return sumFish
}

func ageFish(age int) int {
	if age == 0 {
		age = 6
		return age
	}

	age--
	return age
}

func getFish(data []string) []int {
	var randomNumbersString = strings.Split(data[0], ",")
		var fish = make([]int, len(randomNumbersString))
		for i, v := range randomNumbersString {
			var intVal, _ = strconv.Atoi(v)
			fish[i] = intVal
		}

		return fish
}
