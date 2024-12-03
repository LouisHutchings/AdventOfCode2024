package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	leftList, rightList := getInputLists()
	hashmap := map[int]int{}
	for i := 0; i < len(leftList); i++ {
		hashmap[leftList[i]] = 0
	}
	fmt.Println("hashmap size:", len(hashmap))
	for i := 0; i < len(rightList); i++ {
		_, ok := hashmap[rightList[i]]
		if ok {
			hashmap[rightList[i]]++
		}
	}

	similarityScore := 0
	for key, value := range hashmap {
		similarityScore += key * value
	}
	fmt.Println("similarityScore", similarityScore)
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
		getNumArray(scanner)

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
