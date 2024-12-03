package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	leftList, rightList := getInputLists()
	slices.Sort(leftList)
	slices.Sort(rightList)
	fmt.Println("Sorted Left List: ", leftList[0:10])
	fmt.Println("Sorted right list: ", rightList[0:10])
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += getDistance(leftList[i], rightList[i])
	}
	fmt.Println("Total Distance: ", totalDistance)
}
func getDistance(left int, right int) int {
	distance := left - right
	if distance < 0 {
		distance *= -1
	}
	return distance
}

func getInputLists() ([]int, []int) {
	file, err := os.Open("/Users/louishut/Documents/adventOfCode/day1input")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var leftList []int
	var rightList []int
	isLeftList := true
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		check(err)
		if isLeftList {
			leftList = append(leftList, x)
		} else {
			rightList = append(rightList, x)
		}
		isLeftList = !isLeftList
	}
	fmt.Println("Left List Length:", len(leftList))
	fmt.Println("Left List First Item:", leftList[0])
	fmt.Println("Right List Length:", len(rightList))
	fmt.Println("Right List First Item:", rightList[0])
	return leftList, rightList
}
