package utils

import (
	"fmt"
	"strconv"
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

func ToInt(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}

func PrintAnswer(part int, answer int) {
	fmt.Printf(" ✨ Part %d answer: \033[35m%v\033[0m \n", part, answer)
}

func Debug(s string) {
	if DebugEnabled {
		fmt.Printf("\033[35m%s\n\033[0m", s)
	}
}
