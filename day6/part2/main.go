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
	guardStartingLocation := getGuardLocation(input)
	wallLocations := getWallLocations(input)
	lm := labMap.New(wallLocations, guardStartingLocation, len(input))
	seenLocations := initSeenLocations(input)
	for position, err := lm.GetGuardPosition(); err == nil; position, err = lm.MoveGuard() {
		seenLocations[position[1]][position[0]] = 1
	}
	fmt.Println(countSeenLocations(seenLocations))
	fmt.Println(findNewObectPlacements(seenLocations, guardStartingLocation, wallLocations))
}

func findNewObectPlacements(seenLocations [][]int, guardStartingLocation []int, wallLocations [][]int) int {
	seenLocations[guardStartingLocation[0]][guardStartingLocation[1]] = 0
	newObjectCount := 0
	for i, row := range seenLocations {
		for j, cell := range row {
			if cell == 1 {
				newWallLocations := append(wallLocations, []int{i, j})
				if isGuardInCycle(newWallLocations, guardStartingLocation, seenLocations) {
					newObjectCount++
				}
			}
		}
	}
	return newObjectCount
}

func isGuardInCycle(newWallLocations [][]int, guardStartingLocation []int, seenLocations [][]int) bool {
	lm := labMap.New(newWallLocations, guardStartingLocation, len(seenLocations))
	seenLocationsWithDirection := initSeenLocationsWithDirection(seenLocations)
	for position, err := lm.GetGuardPosition(); err == nil; position, err = lm.MoveGuard() {
		alreadySeenLocationWithDirection := seenLocationsWithDirection[position[1]][position[0]][position[2]] == 1
		if alreadySeenLocationWithDirection {
			return true
		} else {
			seenLocationsWithDirection[position[1]][position[0]][position[2]] = 1
		}
	}
	return false
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

func initSeenLocationsWithDirection(seenLocations [][]int) [][][]int {
	seenLocationsWithDirection := make([][][]int, len(seenLocations))
	for i, _ := range seenLocationsWithDirection {
		seenLocationsWithDirection[i] = make([][]int, len(seenLocations[0]))
		for j, _ := range seenLocations[i] {
			seenLocationsWithDirection[i][j] = make([]int, 4)
		}
	}
	return seenLocationsWithDirection
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
