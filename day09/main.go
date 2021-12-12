package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"fmt"
	"log"
	"sort"
	"strings"
)

var (
	HEIGHTMAP [][]int
	HEIGHTMAP_BASIN [][]Point
	WALL = 9
	FLOODFILLCOUNTER = 0
)

func main() {
		log.SetPrefix("day9: ")
    log.SetFlags(0)

		input, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }

		fmt.Println("Day 9")
		utils.PrintAnswer(1, part1(input))
		utils.PrintAnswer(2, part2(input))
}

func part1(data []string) int {
	HEIGHTMAP = utils.Create2DIntArray(data)
	sumRisk := 0

	for row, rows := range HEIGHTMAP {
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
	HEIGHTMAP_BASIN = Create2DPointArray(data)
	lowPoints := [][]int{}

	for row, rows := range HEIGHTMAP_BASIN {
		for col, point := range rows {
			if(
				getPoint(row-1, col) > point.val &&
				getPoint(row+1, col) > point.val &&
				getPoint(row, col-1) > point.val &&
				getPoint(row, col+1) > point.val) {
					lowPoints = append(lowPoints, []int{row, col})
			}
		}
	}

	// Loop through and fill
	basin_sizes := []int{}
	for _, lowPoint := range lowPoints {
		// fmt.Println(lowPoint)
		FLOODFILLCOUNTER = 0
		flood_fill(lowPoint[0], lowPoint[1])
		basin_sizes = append(basin_sizes, FLOODFILLCOUNTER)
	}

	// fmt.Println(basin_sizes)
	sort.Ints(basin_sizes[:])
	// fmt.Println(basin_sizes)

	return basin_sizes[len(basin_sizes)-1] * basin_sizes[len(basin_sizes)-2] * basin_sizes[len(basin_sizes)-3]
}

func riskLevel(point int) int {
	return point + 1
}

func getPoint(row, col int) int {
	if row < 0 || col < 0 {
		return 9
	}

	if row > len(HEIGHTMAP) -1 || col > len(HEIGHTMAP[row]) -1 {
		return 9
	}

	return HEIGHTMAP[row][col]
}

func getPointPoint(row, col int) Point {
	if row < 0 || col < 0 {
		return Point { 9, false }
	}

	if row > len(HEIGHTMAP_BASIN) -1 || col > len(HEIGHTMAP_BASIN[row]) -1 {
		return Point { 9, false }
	}

	return HEIGHTMAP_BASIN[row][col]
}
type Point struct {
	val int
	basin bool
}

func flood_fill(pos_x, pos_y int) {

   if(getPointPoint(pos_x, pos_y).val == WALL || getPointPoint(pos_x, pos_y).basin) {
		 return
	 }

   HEIGHTMAP_BASIN[pos_x][pos_y].basin = true; // mark the point so that I know if I passed through it.
	 FLOODFILLCOUNTER++

   flood_fill(pos_x + 1, pos_y);  // then i can either go south
   flood_fill(pos_x - 1, pos_y);  // or north
   flood_fill(pos_x, pos_y + 1);  // or east
   flood_fill(pos_x, pos_y - 1);  // or west
}

func Create2DPointArray(data []string) [][]Point {
    var newArr = make([][]Point, len(data))

	for row, line := range data {
		lineArr := strings.Split(line, "")
		newArr[row] = make([]Point, len(lineArr))
		// fmt.Println(lineArr)
		for col, height := range lineArr {
			// fmt.Println(height)
			newArr[row][col] = Point { utils.ToInt(height), false }
		}
	}

    return newArr
}
