package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	FilePath  = "/Users/louishut/Documents/adventOfCode/day4input"
	TestPath  = "/Users/louishut/Documents/adventOfCode/day4TestInput"
	TestPath2 = "/Users/louishut/Documents/adventOfCode/day4TestInput2"
)

func main() {
	input := getInput()
	rows := len(input)
	cols := len(input[0])
	total := 0
	for i, line := range input {
		if i-1 < 0 || i+1 >= rows {
			continue
		}
		for j, char := range line {
			if j-1 < 0 || j+1 >= cols {
				continue
			}
			if char == 'A' && isXmas(i, j, input) {
				total++
			}
		}
	}
	fmt.Println(total)
}

func isXmas(i, j int, input []string) bool {
	return isMas(input[i-1][j-1], input[i+1][j+1]) && isMas(input[i-1][j+1], input[i+1][j-1])
}

func isMas(char1, char2 byte) bool {
	if char1 == 'M' && char2 == 'S' {
		return true
	}
	if char1 == 'S' && char2 == 'M' {
		return true
	}
	return false
}

func getInput() []string {
	file, err := os.Open(FilePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("rows: ", len(input), "cols: ", len(input[0]))
	return input
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
