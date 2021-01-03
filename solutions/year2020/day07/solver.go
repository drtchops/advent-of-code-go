package day07

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

var MY_BAG = "shiny gold"

type containedBag struct {
	Color string
	Count int64
}

type bag struct {
	Color    string
	Children []containedBag
}

func parse(input string) map[string]bag {
	lines := strings.Split(input, "\n")
	bags := make(map[string]bag)
	colorRe := regexp.MustCompile(`^\w+ \w+`)
	childRe := regexp.MustCompile(`(\d+) (\w+ \w+)`)

	for _, line := range lines {
		color := colorRe.FindString(line)
		childMatches := childRe.FindAllStringSubmatch(line, -1)
		children := make([]containedBag, len(childMatches))

		for i, matches := range childMatches {
			count, _ := strconv.ParseInt(matches[1], 10, 64)
			children[i] = containedBag{
				Color: matches[2],
				Count: count,
			}
		}

		bags[color] = bag{
			Color:    color,
			Children: children,
		}
	}

	return bags
}

var bagCache = make(map[string]bool)

func findMyBag(bags map[string]bag, color string) bool {
	if cached, ok := bagCache[color]; ok {
		return cached
	}

	b := bags[color]
	for _, c := range b.Children {
		if c.Color == MY_BAG || findMyBag(bags, c.Color) {
			bagCache[color] = true
			return true
		}
	}

	bagCache[color] = false
	return false
}

func countContents(bags map[string]bag, color string) int64 {
	var count int64
	b := bags[color]
	for _, c := range b.Children {
		count += c.Count + (c.Count * countContents(bags, c.Color))
	}
	return count
}

func (s *Solver) SolveA(input string) string {
	bags := parse(input)
	var count int64

	for _, b := range bags {
		if findMyBag(bags, b.Color) {
			count++
		}
	}

	return fmt.Sprint(count)
}

func (s *Solver) SolveB(input string) string {
	bags := parse(input)
	count := countContents(bags, MY_BAG)
	return fmt.Sprint(count)
}
