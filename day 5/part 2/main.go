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
	invalidPageLists := make([]*PageList, 0)
	for _, pageList := range pageLists {
		isValid := pageList.isValidList(ruleSet)
		if !isValid {
			invalidPageLists = append(invalidPageLists, pageList)
		}
	}
	total := 0
	for _, invalidPageList := range invalidPageLists {
		invalidPageList.fixList(ruleSet)
		increment, err := strconv.Atoi(invalidPageList.getMiddlePage())
		check(err)
		total += increment
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

func (pl *PageList) getMiddlePage() string {
	return pl.pages[len(pl.pages)/2]
}

func (pl *PageList) fixList(rs *RuleSet) {
	badPages := pl.extractBadPages(rs)
	for _, page := range badPages {
		pl.insertPage(page, rs)
	}
}

func (pl *PageList) insertPage(newPage string, rs *RuleSet) {
	pagesMustBeSeen := make(map[string]bool)
	for _, rulePage := range rs.pageMap[newPage] {
		for _, page := range pl.pages {
			if rulePage == page {
				pagesMustBeSeen[page] = true
			}
		}
	}
	newIndex := len(pl.pages) - 1
	for i := newIndex; len(pagesMustBeSeen) > 0; i-- {
		if pagesMustBeSeen[pl.pages[i]] {
			delete(pagesMustBeSeen, pl.pages[i])
		}
		newIndex = i
	}
	pl.pages = append(pl.pages[:newIndex+1], pl.pages[newIndex:]...)
	pl.pages[newIndex] = newPage

}

func (pl *PageList) extractBadPages(rs *RuleSet) []string {
	seenPages := make(map[string]struct{})
	badPages := make([]string, 0)
	goodPages := make([]string, len(pl.pages))
	copy(goodPages, pl.pages)
	for i, page := range pl.pages {
		pagesAfter := rs.pageMap[page]
		for _, pageAfter := range pagesAfter {
			if _, ok := seenPages[pageAfter]; ok {
				badPageIndex := i - len(badPages)
				goodPages = append(goodPages[:badPageIndex], goodPages[badPageIndex+1:]...)
				badPages = append(badPages, page)
				break
			}
		}
		seenPages[page] = struct{}{}
	}
	pl.pages = goodPages
	return badPages
}

func (pl *PageList) isValidList(rs *RuleSet) bool {
	seenPages := make(map[string]struct{})
	for _, page := range pl.pages {
		pagesAfter := rs.pageMap[page]
		for _, pageAfter := range pagesAfter {
			if _, ok := seenPages[pageAfter]; ok {
				return false
			}
		}
		seenPages[page] = struct{}{}
	}
	return true
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
