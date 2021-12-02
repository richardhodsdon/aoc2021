package main

import (
	"aoc/fileloading"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
		log.SetPrefix("day2: ")
    log.SetFlags(0)

		commands, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }

		// Day 2 - Part 1
		var hor, depth1, depth2, aim int = 0, 0, 0, 0
		for _, line := range commands {
			directions := strings.Fields(line)
			value, _ := strconv.Atoi(directions[1])

			switch directions[0] {
				case "forward":
					hor += value
					depth2 += value * aim
				case "down":
					depth1 += value
					aim += value
				case "up":
					depth1 -= value
					aim -= value
			}
		}

		fmt.Println(hor * depth1)
		fmt.Println(hor * depth2)
}
