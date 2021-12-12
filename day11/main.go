package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"fmt"
	"log"
	"strings"
)

var (
	octopii [][]Point
	flashes int = 0
)

func main() {
		log.SetPrefix("day11: ")
    log.SetFlags(0)

		input, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }

		fmt.Println("Day 11")
		utils.PrintAnswer(1, part1(input))
		utils.PrintAnswer(2, part2(input))
}

func part1(data []string) int {
	flashes = 0
	octopii = Create2DPointArray(data)

	// daily increase
	for i := 0; i < 100; i++ {
		resetFlashes()
		for col, rows := range octopii {
			for row, octopus := range rows {
				if octopus.flashed {
					continue
				}

				incr(octopii[col][row])
			}
		}

		// fmt.Println(i+1)
		// printOcto()
	}

	return flashes
}

func part2(data []string) int {
	flashes = 0
	octopii = Create2DPointArray(data)
	var currentStep int

	// daily increase
	for i := 0; i < 1000; i++ {
		currentStep = i
		syncedFlashes := flashes
		resetFlashes()
		for col, rows := range octopii {
			for row, octopus := range rows {
				if octopus.flashed {
					continue
				}

				incr(octopii[col][row])
			}
		}

		if flashes - syncedFlashes == 100 {
			// printOcto()
			break
		}

		// fmt.Println(i+1)
		// printOcto()
	}

	return currentStep + 1
}

type Point struct {
	x, y, value int
	flashed bool
}

func Create2DPointArray(data []string) [][]Point {
    var newArr = make([][]Point, len(data))

	for row, line := range data {
		lineArr := strings.Split(line, "")
		newArr[row] = make([]Point, len(lineArr))
		for col, value := range lineArr {
			newArr[row][col] = Point{ row, col, utils.ToInt(value), false }
		}
	}

    return newArr
}

func flash(octopus Point) {
	x := octopus.x
	y := octopus.y

	if octopus.flashed {
		return
	}

	flashes++

	octopii[x][y].value = 0
	octopii[x][y].flashed = true

	// increase all points
	incr(getPoint(x-1, y-1))
	incr(getPoint(x-1, y))
	incr(getPoint(x-1, y+1))
	incr(getPoint(x, y-1))
	incr(getPoint(x, y+1))
	incr(getPoint(x+1, y-1))
	incr(getPoint(x+1, y))
	incr(getPoint(x+1, y+1))
}

func incr(octopus Point) {
	x := octopus.x
	y := octopus.y

	if octopus.flashed {
		return
	}

	octopii[x][y].value++

	if octopii[x][y].value > 9 {
		flash(octopii[x][y])
	}
}

func resetFlashes() {
	for col, rows := range octopii {
		for row := range rows {
			octopii[col][row].flashed = false
		}
	}
}

func printOcto() {
	for _, rows := range octopii {
		for _, octopus := range rows {
			fmt.Print(octopus.value)
		}
		fmt.Println()
	}
}

func getPoint(row, col int) Point {
	if row < 0 || col < 0 {
		return Point {0,0,0, true}
	}

	if row > len(octopii) -1 || col > len(octopii[row]) -1 {
		return Point {0,0,0, true}
	}

	return octopii[row][col]
}
