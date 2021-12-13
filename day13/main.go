package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var ()

func main() {
		input, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }

		fmt.Println(utils.GetDay())
		utils.PrintAnswer(1, part1(input))
		utils.PrintAnswer(2, part2(input))
}

func part1(data []string) int {
	positions := []string{}
	folds := []string{}
	for _, v := range data {
		if strings.Contains(v, ",") {
			positions = append(positions, v)
		}
		if strings.Contains(v, "=") {
			fold := strings.Split(v, " along ")
			folds = append(folds, fold[1])
		}
	}

	// fmt.Println(positions)

	for i, fold := range folds {
		if i > 0 {
			break
		}
		split := strings.Split(fold, "=")
		// horizontal
		if split[0] == "y" {
			for i, pos := range positions {
				posSplit := strings.Split(pos, ",")

				if utils.ToInt(posSplit[1]) < utils.ToInt(split[1]) {
					continue
				}

				posSplit[1] = strconv.Itoa(utils.ToInt(split[1]) - (utils.ToInt(posSplit[1]) - utils.ToInt(split[1])))
				positions[i] = strings.Join(posSplit, ",")
			}
		}

		if split[0] == "x" {
			for i, pos := range positions {
				posSplit := strings.Split(pos, ",")

				if utils.ToInt(posSplit[0]) < utils.ToInt(split[1]) {
					continue
				}

				posSplit[0] = strconv.Itoa(utils.ToInt(split[1]) - (utils.ToInt(posSplit[0]) - utils.ToInt(split[1])))
				positions[i] = strings.Join(posSplit, ",")
			}
		}

		positions = utils.UniqueString(positions)
	}

	// fmt.Println(positions)
	// fmt.Println(folds)
	return len(positions)
}

func part2(data []string) int {
	positions := []string{}
	folds := []string{}
	for _, v := range data {
		if strings.Contains(v, ",") {
			positions = append(positions, v)
		}
		if strings.Contains(v, "=") {
			fold := strings.Split(v, " along ")
			folds = append(folds, fold[1])
		}
	}

	// fmt.Println(positions)

	for _, fold := range folds {
		split := strings.Split(fold, "=")
		// horizontal
		if split[0] == "y" {
			for i, pos := range positions {
				posSplit := strings.Split(pos, ",")

				if utils.ToInt(posSplit[1]) < utils.ToInt(split[1]) {
					continue
				}

				posSplit[1] = strconv.Itoa(utils.ToInt(split[1]) - (utils.ToInt(posSplit[1]) - utils.ToInt(split[1])))
				positions[i] = strings.Join(posSplit, ",")
			}
		}

		if split[0] == "x" {
			for i, pos := range positions {
				posSplit := strings.Split(pos, ",")

				if utils.ToInt(posSplit[0]) < utils.ToInt(split[1]) {
					continue
				}

				posSplit[0] = strconv.Itoa(utils.ToInt(split[1]) - (utils.ToInt(posSplit[0]) - utils.ToInt(split[1])))
				positions[i] = strings.Join(posSplit, ",")
			}
		}

		positions = utils.UniqueString(positions)
	}

	// fmt.Println(positions)
	// fmt.Println(folds)
	draw(positions)
	return len(positions)
}

func draw(data []string) {
	grid := [6][40]string{}
	for _, v := range data {
		co := strings.Split(v, ",")

		grid[utils.ToInt(co[1])][utils.ToInt(co[0])] = "0"
	}

	// fmt.Println(grid)

	for _, rows := range grid {
		for _, item := range rows {
			if item != "" {
				fmt.Print(item)
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
}
