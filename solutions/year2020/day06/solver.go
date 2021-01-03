package day06

import (
	"fmt"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func keys(dict map[rune]bool) []rune {
	keys := make([]rune, len(dict))
	i := 0
	for k := range dict {
		keys[i] = k
		i++
	}
	return keys
}

func (s *Solver) SolveA(input string) string {
	lines := strings.Split(input, "\n")
	groups := make([][]rune, 0)
	answers := make(map[rune]bool)

	for _, line := range lines {
		if line == "" {
			groups = append(groups, keys(answers))
			answers = make(map[rune]bool)
			continue
		}

		for _, r := range line {
			answers[r] = true
		}
	}
	groups = append(groups, keys(answers))

	var sum int64
	for _, answers := range groups {
		sum += int64(len(answers))
	}
	return fmt.Sprint(sum)
}

func checkGroup(answers map[rune]int, people int) int64 {
	var matches int64
	for _, p := range answers {
		if p == people {
			matches++
		}
	}
	return matches
}

func (s *Solver) SolveB(input string) string {
	var sum int64
	answers := make(map[rune]int)
	people := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			sum += checkGroup(answers, people)
			answers = make(map[rune]int)
			people = 0
			continue
		}

		people++
		for _, r := range line {
			answers[r]++
		}
	}
	sum += checkGroup(answers, people)

	return fmt.Sprint(sum)
}
