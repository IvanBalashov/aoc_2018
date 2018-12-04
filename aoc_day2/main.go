package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	path := "./boxids"
	firstStar(path)
	secondStar(path)
}

func firstStar(path string) {
	fmt.Printf("Advent of code day 2 star one\n")
	boxIds := readFile(path)
	storage := []map[int]int{}

	for i := 0; i < len(boxIds); i++ {
		var buf = map[int]int{}
		var tmp = make(map[int32]int, 0)
		for _, v := range boxIds[i] {
			if _, ok := tmp[v]; ok {
				tmp[v]++
			} else {
				tmp[v] = 1
			}
		}
		for _, value := range tmp {
			if _, ok := buf[value]; !ok {
				buf[value] = 1
			}
		}
		storage = append(storage, buf)
	}
	ar := [26]int{}
	for i := range storage {
		for k := range storage[i] {
			if k <= 1 {
				continue
			}
			ar[k]++
		}
	}

	var answer = 1
	for i := 2; i < len(ar); i++ {
		if ar[i] == 0 {
			continue
		}
		answer = answer * ar[i]
	}
	fmt.Printf("Answer - %s\n", answer)
}

func secondStar(path string) {
	fmt.Printf("Advent of code day 2 star two\n")
	boxIds := readFile(path)
	for k, v := range boxIds {
		for i := range boxIds[k:] {
			if length, pos := findFunc(v, boxIds[i]); length == 1 {
				answer := v[0:pos[0]] + v[pos[0]+1:]
				fmt.Printf("answer - %s\n", answer)
				return
			}
		}
	}

}

func findFunc(first, second string) (int, []int) {
	var positions = []int{}
	length := 0
	for i := range first {
		if first[i] != second[i] {
			length += 1
			positions = append(positions, i)
		}
	}
	return length, positions
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
