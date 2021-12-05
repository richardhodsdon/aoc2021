package main

import (
	"aoc/fileloading"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
		log.SetPrefix("day4: ")
    log.SetFlags(0)

		bingo, err := fileloading.LoadFileAsStringSlice("input.txt")
    if err != nil {
        log.Fatal(err)
    }

		part1(bingo)
}

func part1(data []string) {

		// Part 1
		var randomNumbersString = strings.Split(data[0], ",")
		var randomNumbers = make([]int, len(randomNumbersString))
		for i, v := range randomNumbersString {
			var intVal, _ = strconv.Atoi(v)
			randomNumbers[i] = intVal
		}
		// fmt.Println(data[0])
		// fmt.Println(randomNumbers)

		var boards = getBoards(data)
		// fmt.Println(len(boards))
		// fmt.Println(boards)
		// fmt.Println(boards[0].winCount(randomNumbers))
		// fmt.Println(boards[0].winCount(randomNumbers))

		var winningBoard Board = boards[0];
		for i := 0; i < len(boards); i++ {
			currentCount, _, _ := boards[i].winCount(randomNumbers)
			winningCount, _, _ := winningBoard.winCount(randomNumbers)
			if currentCount > winningCount {
				winningBoard = boards[i]
			}
		}

		_, lastWinningNumber, score := winningBoard.winCount(randomNumbers)
		// fmt.Println(winningBoard.numbers)
		fmt.Println(score)
		fmt.Println(lastWinningNumber)
		// fmt.Println(winningBoard.score(lastWinningNumber))
}

type Board struct {
	numbers [5][5]int
	matchedNumbers [5][5]int
	winCounted int
	winningNumber int
	finalScore int
}

func (v Board) winCount(randomNumbers []int) (int, int, int) {
	if v.winCounted != 0 {
		return v.winCounted, v.winningNumber, v.finalScore
	}

	var count = 0
	v.matchedNumbers = v.numbers
	// fmt.Println(randomNumbers)
	out:
	for _, num := range randomNumbers {
		for rowI, rows := range v.numbers {
			for colI, val := range rows {
				if num == val {
					v.matchedNumbers[rowI][colI] = 999
					// fmt.Print(num)
					// fmt.Print(":")
					// fmt.Println(val)
					// fmt.Println(rowI)
					// fmt.Println(colI)
					// fmt.Println(v.matchedNumbers)

					if v.checkWin() {
						// fmt.Println(v.matchedNumbers)
						// fmt.Println("checkWin")

						v.winCounted = count
						v.winningNumber = num
						break out
					}
				}
			}
		}
		count++
	}

	return v.winCounted, v.winningNumber, v.score(v.winningNumber)
}

func (v Board) checkWin() bool {
	var winning = false
	// row check
	for _, row := range v.matchedNumbers {
		if row == [5]int{999, 999, 999, 999, 999} {
			winning = true
			break
		}
	}

	if winning {
		return winning
	}

	// column check
	for i := 0; i < 5; i++ {
		if (
			v.matchedNumbers[0][i] == 999 &&
			v.matchedNumbers[1][i] == 999 &&
			v.matchedNumbers[2][i] == 999 &&
			v.matchedNumbers[3][i] == 999 &&
			v.matchedNumbers[4][i] == 999) {
				// fmt.Println("vertical win")
				winning = true
				break
		}
	}

	return winning
}

func (v Board) score(lastRandomNumber int) int {
	var total = 0;
	// fmt.Println(v.matchedNumbers)
	for _, rows := range v.matchedNumbers {
		for _, val := range rows {
			if val != 999 {
				total += val
			}
		}
	}
	// fmt.Println(total)
	// fmt.Println(lastRandomNumber)
	return total * lastRandomNumber
}

func getBoards(data []string) []Board {
	var boardIndex = 0
	var boards []Board
	var board Board
	for index, line := range data {
			// skip random numbers and first blank line
			if index < 2 {
				continue
			}

			if line == "" {
				boardIndex = 0
				if (len(board.numbers) > 1) {
					boards = append(boards, board)
				}
				board = Board{}
				continue
			}

			boardLine := strings.Fields(line)

			for lineIndex := 0; lineIndex < len(boardLine); lineIndex++ {
				num, _ := strconv.Atoi(boardLine[lineIndex])
				board.numbers[boardIndex][lineIndex] = num
			}

			boardIndex++
		}

	return boards
}
