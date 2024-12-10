package main

import (
	"bufio"
	"day6/labMap"
	"fmt"
	"os"
)

const (
	FilePath = "/Users/louishut/Documents/adventOfCode/day6input"
	TestPath = "/Users/louishut/Documents/adventOfCode/day6TestInput"
)

func main() {
	input := getInput()
	lm := labMap.New(getWallLocations(input), getGuardLocation(input), len(input))
	seenLocations := initSeenLocations(input)
	for position, err := lm.GetGuardPosition(); err == nil; position, err = lm.MoveGuard() {
		seenLocations[position[1]][position[0]] = 1
	}
	fmt.Println(countSeenLocations(seenLocations))
}

func countSeenLocations(seenLocations [][]int) int {
	total := 0
	for _, row := range seenLocations {
		for _, location := range row {
			total += location
		}
	}
	return total
}
func getInput() []string {
	file, err := os.Open(FilePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	fmt.Println("rows: ", len(input), "cols: ", len(input[0]))
	return input
}

func initSeenLocations(input []string) [][]int {
	seenLocations := make([][]int, len(input))
	for i, _ := range seenLocations {
		seenLocations[i] = make([]int, len(input[0]))
	}
	return seenLocations
}

func getWallLocations(input []string) [][]int {
	var wallLocations [][]int
	for i, line := range input {
		for j, char := range line {
			if char == '#' {
				wallLocations = append(wallLocations, []int{i, j})
			}
		}
	}
	return wallLocations
}

func getGuardLocation(input []string) []int {
	for i, line := range input {
		for j, char := range line {
			if char == '^' {
				return []int{j, i}
			}
		}
	}
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
