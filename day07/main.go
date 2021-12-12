package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"fmt"
	"log"
	"math"
)

func main() {
		log.SetPrefix("day7: ")
    log.SetFlags(0)

		crabs, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }

		fmt.Println("Day 7")
		utils.PrintAnswer(1, part1(crabs))
		utils.PrintAnswer(2, part2(crabs))
}

func part1(data []string) int {
	var crabs = utils.GetIntArrayFromCSV(data)
	var crabMap = make(map[int]map[string]int, 0)

	min, max := utils.MinMax(crabs)

	// fmt.Println(min)
	// fmt.Println(max)

	for i := min; i <= max; i++ {
		_, ok := crabMap[i]
    if !ok {
        mm := make(map[string]int)
        crabMap[i] = mm
    }
	}

	// Create initial HashMap
	for _, crab := range crabs {
		crabMap[crab]["sum"] = 0
		crabMap[crab]["qty"] += 1
	}
	// fmt.Println(crabMap)

	// Work out Costs of Fuel
	var minFuel, maxFuel int
	var initialFlag = false
	for crabHorPos := range crabMap {
		for ccHorPos, countingCrab := range crabMap {
			crabMap[crabHorPos]["sum"] += int(math.Abs(float64(crabHorPos - ccHorPos))) * countingCrab["qty"]
		}

		if !initialFlag {
			initialFlag = true
			minFuel = crabMap[crabHorPos]["sum"]
			maxFuel = crabMap[crabHorPos]["sum"]
		}

		// fmt.Println(crabMap[crabHorPos]["sum"] )
		if maxFuel < crabMap[crabHorPos]["sum"] {
				maxFuel = crabMap[crabHorPos]["sum"]
		}
		if minFuel > crabMap[crabHorPos]["sum"] {
				minFuel = crabMap[crabHorPos]["sum"]
		}
	}

	// fmt.Println(crabMap)
	// fmt.Println(minFuel)
	// fmt.Println(maxFuel)

	return minFuel
}

func part2(data []string) int {
var crabs = utils.GetIntArrayFromCSV(data)
	var crabMap = make(map[int]map[string]int, 0)

	min, max := utils.MinMax(crabs)

	// fmt.Println(min)
	// fmt.Println(max)

	for i := min; i <= max; i++ {
		_, ok := crabMap[i]
    if !ok {
        mm := make(map[string]int)
        crabMap[i] = mm
    }
	}

	// Create initial HashMap
	for _, crab := range crabs {
		crabMap[crab]["sum"] = 0
		crabMap[crab]["qty"] += 1
	}
	// fmt.Println(crabMap)

	// Work out Costs of Fuel
	var minFuel, maxFuel int
	var initialFlag = false
	for crabHorPos := range crabMap {
		for ccHorPos, countingCrab := range crabMap {
			crabMap[crabHorPos]["sum"] += incrFuelCost(crabHorPos, ccHorPos) * countingCrab["qty"]
		}

		if !initialFlag {
			initialFlag = true
			minFuel = crabMap[crabHorPos]["sum"]
			maxFuel = crabMap[crabHorPos]["sum"]
		}

		// fmt.Println(crabMap[crabHorPos]["sum"] )
		if maxFuel < crabMap[crabHorPos]["sum"] {
				maxFuel = crabMap[crabHorPos]["sum"]
		}
		if minFuel > crabMap[crabHorPos]["sum"] {
				minFuel = crabMap[crabHorPos]["sum"]
		}
	}

	// fmt.Println(crabMap)
	// fmt.Println(minFuel)
	// fmt.Println(maxFuel)

	return minFuel
}

func incrFuelCost(start int, end int) int {
	var diff = int(math.Abs(float64(start - end)))

	sum := 0
	for i := 0; i <= diff; i++ {
		sum += i
	}
	return sum
}
