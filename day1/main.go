package main

import (
	"aoc/fileloading"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("day1: ")
    log.SetFlags(0)

		depths, err := fileloading.LoadFileAsIntSlice("input")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

		// Day 1
		var count, countWindow, prevValue int = 0, 0, 0
		for i, value := range depths {
			// currentValue, _ := strconv.Atoi(value)
			if (i == 0 || i >= len(depths) - 2) {
				continue
			}

			if (value > prevValue) {
				count++
			}

			if (depths[i] + depths[i + 1] + depths[i + 2] >	depths[i - 1] + depths[i] + depths[i + 1]) {
				countWindow++;
			}

			prevValue = value
		}

		fmt.Println(count)
		fmt.Println(countWindow)
}
