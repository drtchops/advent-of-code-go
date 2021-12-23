package day10

import (
	"fmt"
	"sort"

	"github.com/drtchops/advent-of-code-go/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func parse(input string) []int64 {
	adapters := utils.ParseInputInts(input, "\n")
	sort.Slice(adapters, func(i, j int) bool { return adapters[i] < adapters[j] })
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	return append([]int64{0}, adapters...)
}

func (s *Solver) SolveA(input string) string {
	adapters := parse(input)

	var last int64
	var count1 int64
	var count3 int64
	for i, a := range adapters {
		if i == 0 {
			continue
		}

		diff := a - last
		if diff == 1 {
			count1++
		} else if diff == 3 {
			count3++
		}
		last = a
	}

	return fmt.Sprint(count1 * count3)
}

func (s *Solver) SolveB(input string) string {
	adapters := parse(input)

	nextSteps := make(map[int][]int, len(adapters))
	for i, a := range adapters {
		max := a + 3
		next := make([]int, 0)
		for j := i + 1; j < len(adapters); j++ {
			if adapters[j] > max {
				break
			}
			next = append(next, j)
		}
		nextSteps[i] = next
	}

	branchCount := make(map[int]int64, len(adapters))
	for i := len(adapters) - 1; i >= 0; i-- {
		next := nextSteps[i]

		if len(next) == 0 {
			branchCount[i] = 1
			continue
		}
		if len(next) == 1 {
			branchCount[i] = branchCount[next[0]]
			continue
		}

		var branches int64
		for _, n := range next {
			branches += branchCount[n]
		}

		branchCount[i] = branches
	}

	return fmt.Sprint(branchCount[0])
}
