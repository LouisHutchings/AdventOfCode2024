package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputLists := getInputLists()
	safeReports := 0
	for _, inputList := range inputLists {
		if isSafeReport(inputList) {
			safeReports++
		}
	}
	fmt.Println("Safe Reports: ", safeReports)
}

func isSafeReport(inputList []int) bool {
	if inputList[0] > inputList[1] {
		return validateDecreasingList(inputList)
	}
	return validateIncreasingList(inputList)
}

func validateIncreasingList(inputList []int) bool {
	for i := 0; i < len(inputList)-1; i++ {
		if inputList[i] >= inputList[i+1] || inputList[i+1]-inputList[i] > 3 {
			return false
		}
	}
	return true
}

func validateDecreasingList(inputList []int) bool {
	for i := 0; i < len(inputList)-1; i++ {
		if inputList[i] <= inputList[i+1] || inputList[i]-inputList[i+1] > 3 {
			return false
		}
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInputLists() [][]int {
	file, err := os.Open("/Users/louishut/Documents/adventOfCode/day2input")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var inputLists [][]int
	for scanner.Scan() {
		inputLists = append(inputLists, getNumArray(scanner))
	}
	fmt.Println("input list 1", inputLists[0])
	fmt.Println("input list 2", inputLists[1])
	fmt.Println("input list 3", inputLists[2])
	return inputLists
}

func getNumArray(scanner *bufio.Scanner) []int {
	numStrings := strings.Split(scanner.Text(), " ")
	nums := make([]int, len(numStrings))
	for i, numString := range numStrings {
		num, err := strconv.Atoi(numString)
		check(err)
		nums[i] = num
	}
	return nums
}
