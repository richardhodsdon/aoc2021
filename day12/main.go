package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"fmt"
	"log"
)

var ()

func main() {
		input, err := fileloading.LoadFileAsStringSlice("input-single.txt")
    if err != nil {
        log.Fatal(err)
    }

		fmt.Println(utils.GetDay())
		utils.PrintAnswer(1, part1(input))
		utils.PrintAnswer(2, part2(input))
}

func part1(data []string) int {

	return 1
}

func part2(data []string) int {

	return 2
}
