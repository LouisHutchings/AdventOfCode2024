package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	FilePath  = "/Users/louishut/Documents/adventOfCode/day4input"
	TestPath  = "/Users/louishut/Documents/adventOfCode/day4TestInput"
	TestPath2 = "/Users/louishut/Documents/adventOfCode/day4TestInput2"
)

func main() {
	input := getInput()
	fmt.Println(countWords(input))
}

func countWords(input []string) int {
	horizontalWords := countWordsHorizontal(input)
	fmt.Println(horizontalWords)
	verticalWords := countWordsVertical(input)
	fmt.Println(verticalWords)
	diagonalWords := countWordsDiagonal(input)
	fmt.Println(diagonalWords)
	return horizontalWords + verticalWords + diagonalWords
}

func countWordsHorizontal(input []string) int {
	totalWords := 0
	for _, line := range input {
		totalWords += strings.Count(line, "XMAS")
		totalWords += strings.Count(line, "SAMX")
	}
	return totalWords
}

func countWordsVertical(input []string) int {
	return countWordsHorizontal(transposeInput(input))
}

func countWordsDiagonal(input []string) int {
	return countWordsHorizontal(diagonalRightTransform(input)) +
		countWordsHorizontal(diagonalLeftTransform(input))
}

func transposeInput(input []string) []string {
	newInput := make([]string, len(input[0]))
	for _, line := range input {
		for j, char := range line {
			newInput[j] = newInput[j] + string(char)
		}
	}
	return newInput
}

func diagonalRightTransform(input []string) []string {
	newInput := make([]string, len(input)*2-1)
	rows := len(input)
	cols := len(input[0])
	for i := 0; i < cols; i++ {
		newRow := ""
		for j := 0; j < rows-i; j++ {
			newRow += string(input[i+j][j])
		}
		newInput[i] = newRow
		newRow = ""
		for j := 0; j < rows-i; j++ {
			newRow += string(input[j][i+j])
		}
		newInput[i+rows-1] = newRow

	}
	return newInput
}

func diagonalLeftTransform(input []string) []string {
	newInput := make([]string, len(input)*2-1)
	rows := len(input)
	cols := len(input[0])
	for i := cols - 1; i >= 0; i-- {
		newRow := ""
		for j := 0; j <= i; j++ {
			newRow += string(input[i-j][j])
		}
		newInput[i] = newRow
		newRow = ""
		for j := 0; j <= i; j++ {
			newRow += string(input[rows-1-j][cols-1-i+j])
		}
		newInput[i+rows-1] = newRow
	}
	return newInput
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
