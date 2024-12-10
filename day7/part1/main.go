package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	FilePath = "/Users/louishut/Documents/adventOfCode/day7input"
	TestPath = "/Users/louishut/Documents/adventOfCode/day7TestInput"
)

func main() {
	input := getInput()
	total := int64(0)
	for _, line := range input {
		splitString := strings.Split(line, ":")
		target, err := strconv.Atoi(splitString[0])
		check(err)
		var numbers []int
		for _, numberString := range strings.Split(strings.TrimSpace(splitString[1]), " ") {
			number, err := strconv.Atoi(numberString)
			check(err)
			numbers = append(numbers, number)
		}
		makesTarget := canMakeTarget(target, numbers)
		if makesTarget {
			total += int64(target)
		}
		fmt.Println("target:", target, "numbers:", numbers, "canMakeTarget:", makesTarget)
	}
	fmt.Println(total)
}

func canMakeTarget(target int, numbers []int) bool {
	var operatorCombinations []string
	combinations := math.Pow(float64(2), float64(len(numbers)-1))
	for i := 0; float64(i) < combinations; i++ {
		operatorCombinations = append(operatorCombinations, strconv.FormatInt(int64(i), 2))
	}
	for i, operatorCombination := range operatorCombinations {
		for len(operatorCombination) < len(numbers)-1 {
			operatorCombination = "0" + operatorCombination
		}
		operatorCombinations[i] = operatorCombination
	}
	for _, operatorCombination := range operatorCombinations {
		if target == processOperatorCombinations(target, numbers, operatorCombination) {
			return true
		}
	}
	return false
}

func processOperatorCombinations(target int, numbers []int, operatorCombination string) int {
	if len(numbers) != len(operatorCombination)+1 {
		panic("not enough operators")
	}
	total := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if total >= target {
			return total
		}
		if operatorCombination[i-1] == '1' {
			total *= numbers[i]
		} else {
			total += numbers[i]
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
	return input
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
