package utils

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var DebugEnabled = true

func GetDay() string {
    dir, _ := os.Getwd()

    return "Day " + dir[len(dir) - 2:]
}

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

func PrintAnswerTime(part int, answer int, startTime time.Time) {
    elapsed := time.Since(startTime)

	fmt.Printf(" ✨ Part %d answer: \033[35m%v\033[0m in \033[34m%s\033[0m \n", part, answer, elapsed)
}

func Debug(s string) {
	if DebugEnabled {
		fmt.Printf("\033[35m%s\n\033[0m", s)
	}
}

func ReverseMap(m map[string]string) map[string]string {
    n := make(map[string]string, len(m))
    for k, v := range m {
        n[v] = k
    }
    return n
}

func UniqueInt(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func UniqueString(stringSlice []string) []string {
    keys := make(map[string]bool)
    list := []string{}
    for _, entry := range stringSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func SortStringIntMap(si map[string]int) []KV {
    var ss []KV
    for k, v := range si {
        ss = append(ss, KV{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })

    return ss
}

func SortRuneInt64Map(si map[rune]int64) []KRV {
    var ss []KRV
    for k, v := range si {
        ss = append(ss, KRV{k, v})
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })

    return ss
}

type KRV struct {
    Key   rune
    Value int64
}

type KV struct {
    Key   string
    Value int
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	out := make(map[K]V, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}
