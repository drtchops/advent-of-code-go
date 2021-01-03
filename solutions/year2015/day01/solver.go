package day01

import (
	"fmt"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolveA(input string) string {
	directions := strings.Split(input, "")
	var floor int64
	for _, dir := range directions {
		if dir == "(" {
			floor += 1
		} else {
			floor -= 1
		}
	}
	return fmt.Sprint(floor)
}

func (s *Solver) SolveB(input string) string {
	directions := strings.Split(input, "")
	floor := 0
	for i, dir := range directions {
		if dir == "(" {
			floor += 1
		} else {
			floor -= 1
		}
		if floor == -1 {
			return fmt.Sprint(int64(i) + 1)
		}
	}
	return ""
}
