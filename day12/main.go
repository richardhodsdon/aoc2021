package main

import (
	"aoc/fileloading"
	"aoc/utils"
	"aoc/set"

	"fmt"
	"log"
	"strings"
	"unicode"
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

	connections := make(map[string]set.Set[string])
	for _, line := range data {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			panic(parts)
		}
		a := parts[0]
		b := parts[1]
		if _, ok := connections[a]; !ok {
			connections[a] = set.Set[string]{}
		}
		if _, ok := connections[b]; !ok {
			connections[b] = set.Set[string]{}
		}
		connections[a].Add(b)
		connections[b].Add(a)
	}

	paths := []Path{{pos: "start"}}
	completePaths := []Path{}

	for len(paths) > 0 {
		newPaths := []Path{}
		for _, path := range paths {
			pos := path.pos
			nexts, ok := connections[pos]
			if !ok {
				continue
			}

			for next := range nexts {
				if next == "start" {
					continue
				}
				count := path.visited[next]
				isSmallCave := IsLower(next)
				if !isSmallCave || count == 0 {
					nextVisited := utils.CopyMap(path.visited)
					nextVisited[pos] += 1
					newPath := Path{
						pos:         next,
						visited:     nextVisited,
						doubleVisit: path.doubleVisit || (isSmallCave && count == 1),
					}
					if next == "end" {
						completePaths = append(completePaths, newPath)
					} else {
						newPaths = append(newPaths, newPath)
					}
				}
			}
		}
		paths = newPaths
	}

	// fmt.Printf("Completed paths: %d\n", len(completePaths))
	return len(completePaths)
}

func part2(data []string) int {
	connections := make(map[string]set.Set[string])
	for _, line := range data {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			panic(parts)
		}
		a := parts[0]
		b := parts[1]
		if _, ok := connections[a]; !ok {
			connections[a] = set.Set[string]{}
		}
		if _, ok := connections[b]; !ok {
			connections[b] = set.Set[string]{}
		}
		connections[a].Add(b)
		connections[b].Add(a)
	}

	paths := []Path{{pos: "start"}}
	completePaths := []Path{}

	for len(paths) > 0 {
		newPaths := []Path{}
		for _, path := range paths {
			pos := path.pos
			nexts, ok := connections[pos]
			if !ok {
				continue
			}

			for next := range nexts {
				if next == "start" {
					continue
				}
				count := path.visited[next]
				isSmallCave := IsLower(next)
				if !isSmallCave || count == 0 || (count == 1 && !path.doubleVisit) {
					nextVisited := utils.CopyMap(path.visited)
					nextVisited[pos] += 1
					newPath := Path{
						pos:         next,
						visited:     nextVisited,
						doubleVisit: path.doubleVisit || (isSmallCave && count == 1),
					}
					if next == "end" {
						completePaths = append(completePaths, newPath)
					} else {
						newPaths = append(newPaths, newPath)
					}
				}
			}
		}
		paths = newPaths
	}

	// fmt.Printf("Completed paths: %d\n", len(completePaths))
	return len(completePaths)
}

type Path struct {
	pos         string
	visited     map[string]int
	doubleVisit bool
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
