package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main () {
	path := "./frequency"
	firstStar(path)
	secondStar(path)
}

func getNumber(number string) int64 {
	intNumber, err := strconv.ParseInt(number[1:],10,64)
	if err != nil {
		panic(err)
	}
	return intNumber
}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error while read file - %s\n")
		return nil
	}

	readedData := bytes.NewBuffer(data).String()
	frequency := strings.Split(readedData, "\n")
	return frequency
}

func firstStar(path string) {
	fmt.Printf("Advent of code. Day 1 star one\n")
	frequency := readFile(path)
	var answer int64 = 0
	for _, v := range frequency {
		number, err := strconv.ParseInt(v[1:],10,64)
		if err != nil {
			return
		}
		switch v[0] {
		case '+':
			answer = answer + number
		case '-':
			answer = answer - number
		}
	}

	fmt.Printf("Answer %v\n", answer)
}

func secondStar(path string) {
	fmt.Printf("Advent of code. Day 1 star two\n")
	frequency := readFile(path)
	var answer int64 = 0
	answersBase := map[int64]bool{}
	for i := 0; i < len(frequency); i++ {
		number := getNumber(frequency[i])
		switch frequency[i][0] {
		case '+':
			answer = answer + number
		case '-':
			answer = answer - number
		}
		if _, ok := answersBase[answer]; ok {
			break
		} else {
			answersBase[answer] = false
		}
		if i == len(frequency) - 1 {
			i = -1
		}
	}
	fmt.Printf("Answer %v\n", answer)
}