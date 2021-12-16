package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"fmt"
	"log"
	"strings"
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
	var current string
	var template = make(map[string]string)
	var charCount = make(map[string]int)
	current = data[0]
	for i := 2; i < len(data); i++ {
		mapping := strings.Split(data[i], " -> ")
		template[mapping[0]] = mapping[1]
	}

	// steps
	for i := 0; i < 10; i++ {
		newString := ""
		for i := 0; i < len(current); i++ {
			newString += string(current[i])

			if (i + 1) < len(current) {
				newString += template[string(current[i]) + string(current[i+1])]
			}
		}

		current = newString
		// fmt.Println(current)
	}

	// fmt.Println(current)
	// fmt.Println(template)

	for _, v := range current {
		charCount[string(v)] += 1
	}
	sortedMap := utils.SortStringIntMap(charCount)

	return sortedMap[0].Value - sortedMap[len(sortedMap)-1].Value
}

func part2(data []string) int {
var current string
	var template = make(map[string]rune)
	current = data[0]
	for i := 2; i < len(data); i++ {
		var pre string
		var post rune
		_, err := fmt.Sscanf(data[i], "%s -> %c", &pre, &post)
		if err != nil {
			panic(data[i])
		}

		template[pre] = post
	}

	// steps
	pairs := TemplateToPairs(current)

	for step := 1; step <= 40; step++ {
		pairs = AdvanceMap(pairs, template)
	}

	counts := GetCharCounts(current, pairs)
	sortedMap := utils.SortRuneInt64Map(counts)

	return int(sortedMap[0].Value - sortedMap[len(sortedMap)-1].Value)
}

func TemplateToPairs(template string) map[string]int64 {
	chars := []rune(template)
	res := make(map[string]int64)
	for i := 1; i < len(chars); i++ {
		two := string([]rune{chars[i-1], chars[i]})
		res[two] += 1
	}
	return res
}

func AdvanceMap(pairCounts map[string]int64, rules map[string]rune) map[string]int64 {
	newCounts := make(map[string]int64)
	for pair, count := range pairCounts {
		post, ok := rules[pair]
		if !ok {
			panic(pair)
		}
		pre := []rune(pair)
		a := pre[0]
		b := pre[1]

		newCounts[string([]rune{a, post})] += count
		newCounts[string([]rune{post, b})] += count
	}
	return newCounts
}

func GetCharCounts(initTemplate string, counts map[string]int64) map[rune]int64 {
	out := make(map[rune]int64)
	for pair, count := range counts {
		chars := []rune(pair)
		if len(chars) != 2 {
			panic(pair)
		}
		b := chars[1]
		out[b] += count
	}

	initRunes := []rune(initTemplate)
	out[initRunes[0]] += 1
	return out
}
