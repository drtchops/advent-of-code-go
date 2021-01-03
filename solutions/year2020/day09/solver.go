package day09

import (
	"fmt"

	"github.com/drtchops/aoc/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

var PREAMBLE_LENGTH = 25

func isSum(values []int64, idx int) bool {
	for i := idx - PREAMBLE_LENGTH; i < idx; i++ {
		for j := i + 1; j < idx; j++ {
			if values[i]+values[j] == values[idx] {
				return true
			}
		}
	}

	return false
}

func min(values []int64) int64 {
	m := values[0]
	for _, n := range values {
		if n < m {
			m = n
		}
	}
	return m
}

func max(values []int64) int64 {
	m := values[0]
	for _, n := range values {
		if n > m {
			m = n
		}
	}
	return m
}

func findInvalid(values []int64) int64 {
	for i := PREAMBLE_LENGTH; i < len(values); i++ {
		if !isSum(values, i) {
			return values[i]
		}
	}

	return 0
}

func findWeakness(values []int64, invalid int64) int64 {
	for i, v1 := range values {
		var sum int64 = v1
		numbers := []int64{v1}
		for j := i + 1; j < len(values); j++ {
			v2 := values[j]
			sum += v2
			numbers = append(numbers, v2)
			if sum == invalid {
				return min(numbers) + max(numbers)
			}
			if sum > invalid {
				break
			}
		}
	}

	return 0
}

func (s *Solver) SolveA(input string) string {
	values := utils.ParseInputInts(input, "\n")
	invalid := findInvalid(values)
	return fmt.Sprint(invalid)
}

func (s *Solver) SolveB(input string) string {
	values := utils.ParseInputInts(input, "\n")
	invalid := findInvalid(values)
	weakness := findWeakness(values, invalid)
	return fmt.Sprint(weakness)
}
