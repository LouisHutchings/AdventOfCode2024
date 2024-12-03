package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FilePath = "/Users/louishut/Documents/adventOfCode/day2input"
const TestPath = "/Users/louishut/Documents/adventOfCode/day2TestInput"

func main() {
	inputLists := getInputLists()
	safeReports := 0
	fmt.Println("Total Reports: ", len(inputLists))
	for _, inputList := range inputLists {
		if isSafeReport(inputList) {
			safeReports++
		}
	}
	fmt.Println("Safe Reports: ", safeReports)
}

func isSafeReport(inputList []int) bool {
	isValid := false
	if validateList(inputList) {
		isValid = true
	}
	for i, _ := range inputList {
		tempList := make([]int, len(inputList))
		copy(tempList, inputList)
		if validateList(remove(tempList, i)) {
			isValid = true
		}
	}
	return isValid
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func validateList(inputList []int) bool {
	isValidAscendingList := true
	isValidDescendingList := true
	for i := 0; i < len(inputList)-1; i++ {
		if invalidIncrease(inputList[i], inputList[i+1]) {
			isValidAscendingList = false
			break
		}
	}
	for i := 0; i < len(inputList)-1; i++ {
		if invalidDecrease(inputList[i], inputList[i+1]) {
			isValidDescendingList = false
			break
		}
	}
	return isValidAscendingList || isValidDescendingList
}

func invalidIncrease(num1, num2 int) bool {
	return num1 >= num2 || num2-num1 > 3
}

func invalidDecrease(num1, num2 int) bool {
	return num1 <= num2 || num1-num2 > 3
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInputLists() [][]int {
	file, err := os.Open(FilePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var inputLists [][]int
	for scanner.Scan() {
		inputLists = append(inputLists, getNumArray(scanner))
	}
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
