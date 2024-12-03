package main

import (
	"fmt"
	"os"
)

const FilePath = "/Users/louishut/Documents/adventOfCode/day3input"

type MulFunc struct {
	num1 int
	num2 int
}

const (
	MulFuncToken      = "mul"
	RightBracketToken = "("
	LeftBracketToken  = ")"
	CommaToken        = ","
	DoToken           = "do()"
	DontToken         = "don't()"
)

func main() {
	input := getInput()
	total := 0
	mulEnabled := true
	for i := 0; i < len(input); i++ {
		if input[i] == 'd' {
			isDoToken, _ := readToken(DoToken, input, i)
			if isDoToken {
				mulEnabled = true
				continue
			}
			isDontToken, _ := readToken(DontToken, input, i)
			if isDontToken {
				mulEnabled = false
				continue
			}
		}
		if !mulEnabled {
			continue
		}
		if input[i] != 'm' {
			continue
		}
		mulFunc, ok := readMulFunc(input, i)
		if !ok {
			continue
		}
		total += mulFunc.num1 * mulFunc.num2
	}
	fmt.Println(total)
}

func readToken(token string, input string, pointer int) (bool, int) {
	if pointer > len(input)-len(token) {
		return false, pointer
	}
	for i := 0; i < len(token); i++ {
		if token[i] != input[pointer+i] {
			return false, pointer + i
		}
	}
	return true, pointer + len(token)

}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func readNumber(input string, pointer int) (int, int, bool) {
	number := 0
	newPointer := pointer
	if isDigit(input[pointer]) {
		number += int(input[pointer]) - '0'
		newPointer++
	} else {
		return number, newPointer, false
	}
	for i := 1; i < 3 && (pointer+i) < len(input); i++ {
		if isDigit(input[pointer+i]) {
			number *= 10
			number += int(input[pointer+i]) - '0'
			newPointer++
		} else {
			return number, newPointer, true
		}
	}
	return number, newPointer, true
}

func readMulFunc(input string, pointer int) (MulFunc, bool) {
	isToken, newPointer := readToken(MulFuncToken, input, pointer)
	if !isToken {
		return MulFunc{}, false
	}
	pointer = newPointer
	isToken, newPointer = readToken(RightBracketToken, input, pointer)
	if !isToken {
		return MulFunc{}, false
	}
	pointer = newPointer
	num1, newPointer, ok := readNumber(input, pointer)
	if !ok {
		return MulFunc{}, false
	}
	pointer = newPointer
	isToken, newPointer = readToken(CommaToken, input, pointer)
	if !isToken {
		return MulFunc{}, false
	}
	pointer = newPointer
	num2, newPointer, ok := readNumber(input, pointer)
	if !ok {
		return MulFunc{}, false
	}
	pointer = newPointer
	isToken, newPointer = readToken(LeftBracketToken, input, pointer)
	if !isToken {
		return MulFunc{}, false
	}
	pointer = newPointer
	return MulFunc{num1: num1, num2: num2}, true

}

func getInput() string {
	input, err := os.ReadFile(FilePath)
	check(err)
	return string(input)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
