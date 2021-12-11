package utils

import (
	"fmt"
	"strconv"
	"strings"
)

var DebugEnabled = true

func Contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func ContainsString(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func ToInt(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}

func GetIntArrayFromCSV(data []string) []int {
	var randomNumbersString = strings.Split(data[0], ",")
    var intArr = make([]int, len(randomNumbersString))
    for i, v := range randomNumbersString {
        var intVal, _ = strconv.Atoi(v)
        intArr[i] = intVal
    }

    return intArr
}

func Create2DIntArray(data []string) [][]int {
    var newArr = make([][]int, len(data))

	for row, line := range data {
		lineArr := strings.Split(line, "")
		newArr[row] = make([]int, len(lineArr))
		// fmt.Println(lineArr)
		for col, height := range lineArr {
			// fmt.Println(height)
			newArr[row][col] = ToInt(height)
		}
	}

    return newArr
}

func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}

func PrintAnswer(part int, answer int) {
	fmt.Printf(" ✨ Part %d answer: \033[35m%v\033[0m \n", part, answer)
}

func Debug(s string) {
	if DebugEnabled {
		fmt.Printf("\033[35m%s\n\033[0m", s)
	}
}
