package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	FilePath = "/Users/louishut/Documents/adventOfCode/day5input"
	TestPath = "/Users/louishut/Documents/adventOfCode/day5TestInput"
)

type RuleSet struct {
	pageMap map[string][]string
}

type PageList struct {
	pages []string
}

func main() {
	ruleSet, pageLists := getInput()
	fmt.Println(ruleSet.pageMap)
	total := 0
	for _, pageList := range pageLists {
		isValid, middlePage := pageList.isValidList(ruleSet)
		if isValid {
			increment, err := strconv.Atoi(middlePage)
			check(err)
			total += increment
		}
		fmt.Println(pageList.pages, isValid, middlePage)
	}
	fmt.Println(total)
}

func getInput() (*RuleSet, []*PageList) {
	file, err := os.Open(FilePath)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var pageLists []*PageList
	ruleSet := NewRuleSet()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			ruleSet.parseRule(line)
		}
		if strings.Contains(line, ",") {
			pageLists = append(pageLists, NewPageList(line))
		}
	}
	return ruleSet, pageLists
}

func NewRuleSet() *RuleSet {
	ruleSet := &RuleSet{}
	ruleSet.pageMap = make(map[string][]string)
	return ruleSet
}

func NewPageList(pages string) *PageList {
	pageList := &PageList{}
	pageList.pages = strings.Split(pages, ",")
	return pageList
}

func (pl *PageList) isValidList(rs *RuleSet) (bool, string) {
	seenPages := make(map[string]struct{})
	for _, page := range pl.pages {
		pagesAfter := rs.pageMap[page]
		for _, pageAfter := range pagesAfter {
			if _, ok := seenPages[pageAfter]; ok {
				return false, ""
			}
		}
		seenPages[page] = struct{}{}
	}
	middlePage := pl.pages[len(pl.pages)/2]
	return true, middlePage
}

func (r *RuleSet) parseRule(rule string) {
	pages := strings.Split(rule, "|")
	if len(pages) != 2 {
		panic("invalid rule format")
	}
	r.pageMap[pages[0]] = append(r.pageMap[pages[0]], pages[1])
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
