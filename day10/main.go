package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"fmt"
	"log"
	"sort"
	"strings"
)

var navHashMap = map[string]string{
    "(": ")",
    "[": "]",
    "{": "}",
    "<": ">",
}

var reverseNavHashMap = utils.ReverseMap(navHashMap)

var syntaxCost = map[string]int{
    ")": 3,
    "]": 57,
    "}": 1197,
    ">": 25137,
}

var syntaxScore = map[string]int{
    "(": 1,
    "[": 2,
    "{": 3,
    "<": 4,
}

func main() {
		log.SetPrefix("day10: ")
    log.SetFlags(0)

		input, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }

		fmt.Println("Day 10")
		utils.PrintAnswer(1, part1(input))
		utils.PrintAnswer(2, part2(input))
}

func part1(data []string) int {
	syntaxCosts := 0

	for _, line := range data {
		lineChars := strings.Split(line, "")
		stack := []string{}
		illegalLinebreak:

		for _, character := range lineChars {
			// if we are opening jsut append
			if _, ok := navHashMap[character]; ok {
				stack = append(stack, character)
				continue
			}

			// if it matches then remove from the end
			if stack[len(stack) - 1] == reverseNavHashMap[character] {
				stack = stack[:len(stack) - 1]
				continue
			}

			// We hit an illegal breakpoint
			syntaxCosts += syntaxCost[character]
			break illegalLinebreak
		}

	}


	return syntaxCosts
}

func part2(data []string) int {
	scores := []int{}

	for _, line := range data {
		lineChars := strings.Split(line, "")
		stack := []string{}
		score := 0
		illegalLinebreak:

		for _, character := range lineChars {
			// if we are opening jsut append
			if _, ok := navHashMap[character]; ok {
				stack = append(stack, character)
				continue
			}

			// if it matches then remove from the end
			if stack[len(stack) - 1] == reverseNavHashMap[character] {
				stack = stack[:len(stack) - 1]
				continue
			}

			// We hit an illegal breakpoint
			stack = []string{}
			break illegalLinebreak
		}

		// fmt.Println(stack)

		// toCloseStack := []string{}
		// get the legal remaining
		for i := len(stack); i > 0; i-- {
			// fmt.Println(stack[i-1])
			score = (5 * score) + syntaxScore[stack[i-1]]

		}
		// fmt.Println(score)
		if score != 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores)
	var answer = int(float32(len(scores)-1)/2 + 0.5)
	return scores[answer]
}
