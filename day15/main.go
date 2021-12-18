package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"fmt"
	"log"
	"time"
)

var ()

func main() {
		input, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }
		startTime := time.Now()


		fmt.Println(utils.GetDay())
		utils.PrintAnswerTime(1, part1(input), startTime)
		utils.PrintAnswerTime(2, part2(input), startTime)
}

func part1(data []string) int {
	return 1
}

func part2(data []string) int {
	return 2
}
