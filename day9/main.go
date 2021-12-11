package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"fmt"
	"log"
)

var heightmap [][]int

func main() {
		log.SetPrefix("day9: ")
    log.SetFlags(0)

		input, err := fileloading.LoadFileAsStringSlice("input-single.txt")
    if err != nil {
        log.Fatal(err)
    }

		fmt.Println("Day 9")
		utils.PrintAnswer(1, part1(input))
		utils.PrintAnswer(2, part2(input))
}

func part1(data []string) int {
	heightmap = utils.Create2DIntArray(data)
	sumRisk := 0

	for row, rows := range heightmap {
		for col, val := range rows {
			if(
				getPoint(row-1, col) > val &&
				getPoint(row+1, col) > val &&
				getPoint(row, col-1) > val &&
				getPoint(row, col+1) > val) {
				sumRisk += riskLevel(val)
			}
		}
	}

	// fmt.Println(heightmap)
	return sumRisk
}

func part2(data []string) int {
	return 2
}

func riskLevel(point int) int {
	return point + 1
}

func getPoint(row, col int) int {
	if row < 0 || col < 0 {
		return 9
	}

	if row > len(heightmap) -1 || col > len(heightmap[row]) -1 {
		return 9
	}

	return heightmap[row][col]
}
