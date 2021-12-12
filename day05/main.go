package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"log"
	"strings"
)

func main() {
		log.SetPrefix("day4: ")
    log.SetFlags(0)

		vents, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }

		utils.PrintAnswer(1, part1(vents))
		utils.PrintAnswer(2, part2(vents))
}

func part1(data []string) int {
	var vents = getVents(data)
	var grid [999][999]int
	// fmt.Println(vents)

	for _, vent := range vents {
		for _, point := range vent.getPoints(false) {
			grid[point.x][point.y]++
		}
	}

	return countGrid(grid)
}

func part2(data []string) int {
	var vents = getVents(data)
	var grid [999][999]int

	for _, vent := range vents {
		for _, point := range vent.getPoints(true) {
			grid[point.x][point.y]++
		}
	}

	return countGrid(grid)
}

type Point struct {
	x int
	y int
}

type Vent struct {
	start Point
	end Point
}

func (v Vent) getPoints(includeDiag bool) []Point {
	var allPoints = make([]Point, 0)

	// horizontal points
	if v.start.x == v.end.x {
		for i := v.start.y; i <= v.end.y; i++ {
			allPoints = append(allPoints, Point{ v.start.x, i})
		}

		for i := v.end.y; i <= v.start.y; i++ {
			allPoints = append(allPoints, Point{ v.start.x, i})
		}

		return allPoints
	}

	// Vertical Points
	if v.start.y == v.end.y {
		for i := v.start.x; i <= v.end.x; i++ {
			allPoints = append(allPoints, Point{ i, v.start.y})
		}

		for i := v.end.x; i <= v.start.x; i++ {
			allPoints = append(allPoints, Point{ i, v.start.y,})
		}

		return allPoints
	}

	if !includeDiag {
		return allPoints
	}

	// Diagonals
	// /
	if v.start.x < v.end.x {
		var count = 0
		for i := v.start.x; i <= v.end.x; i++ {
			var newY int
			if v.start.y < v.end.y {
				newY = v.start.y + count
			}	else {
				newY = v.start.y - count
			}
			allPoints = append(allPoints, Point{ i, newY})
			count++
		}
	}

	// \
	if v.start.x > v.end.x {
		var count = 0
		for i := v.start.x; i >= v.end.x; i-- {
			var newY int
			if v.start.y < v.end.y {
				newY = v.start.y + count
			}	else {
				newY = v.start.y - count
			}
			allPoints = append(allPoints, Point{ i, newY })
			count++
		}
	}

	// fmt.Println(allPoints)
	return allPoints
}

func getVents(data []string) []Vent {
	var vents = make([]Vent, 0)
	for _, vs := range data {
		vents = append(vents, parseVentString(vs))
	}

	return vents
}

func parseVentString(ventString string) Vent {
	var newVent Vent
	var parsed = strings.Split(ventString, "->")

	var startCoords = strings.Split(strings.TrimSpace(parsed[0]), ",")
	var endCoords = strings.Split(strings.TrimSpace(parsed[1]), ",")
	newVent.start = Point{ utils.ToInt(startCoords[0]), utils.ToInt(startCoords[1]) }
	newVent.end = Point{ utils.ToInt(endCoords[0]), utils.ToInt(endCoords[1]) }

	return newVent
}

func countGrid(grid [999][999]int) int {
	var count = 0;
	for _, row := range grid {
		for _, val := range row {
			if (val >= 2) {
				count++
			}
		}
	}

	return count
}

