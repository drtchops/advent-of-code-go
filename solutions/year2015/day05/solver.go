package day05

import (
	"fmt"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolveA(input string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	words := strings.Split(input, "\n")
	var nice int64

	for _, word := range words {
		if strings.Contains(word, "ab") || strings.Contains(word, "cd") || strings.Contains(word, "pq") || strings.Contains(word, "xy") {
			continue
		}

		vowelCount := 0
		last := 'z'
		hasDouble := false
		for i, c := range word {
			for _, v := range vowels {
				if c == v {
					vowelCount++
					break
				}
			}
			if i > 0 && c == last {
				hasDouble = true
			}
			last = c
		}

		if hasDouble && vowelCount >= 3 {
			nice++
		}
	}

	return fmt.Sprint(nice)
}

func (s *Solver) SolveB(input string) string {
	words := strings.Split(input, "\n")
	var nice int64

	for _, word := range words {
		last2 := 'z'
		last := 'z'
		hasTriple := false
		hasRepeat := false

		for i, c := range word {
			if !hasRepeat && i > 0 {
				pair := string([]rune{last, c})
				pairIdx := strings.LastIndex(word, pair)
				if pairIdx > i {
					hasRepeat = true
				}
			}
			if !hasTriple && i > 1 && c == last2 {
				hasTriple = true
			}

			if hasTriple && hasRepeat {
				nice++
				break
			}

			last2 = last
			last = c
		}
	}

	return fmt.Sprint(nice)
}
