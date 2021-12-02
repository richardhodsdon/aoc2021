package fileloading

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func LoadFile(filename string) ([]int, error) {
    // data, err := ioutil.ReadFile(filename)
    // if err != nil {
		// 		// If no name was given, return an error with a message.
		// 		return nil, errors.New("file reading error")
    // }

		f, err := os.Open(filename)
    if err != nil {
				return nil, errors.New("file reading error")
    }
    defer func() {
        if err = f.Close(); err != nil {
        log.Fatal(err)
    }
    }()

		s := bufio.NewScanner(f)
		var data []int
    for s.Scan() {
				value, _ := strconv.Atoi(s.Text())

        data = append(data, value)
    }

		return data, nil
}
