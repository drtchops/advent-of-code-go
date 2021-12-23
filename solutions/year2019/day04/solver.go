package day04

import (
	"fmt"
	"strconv"

	"github.com/drtchops/advent-of-code-go/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func isValid(pass int64) bool {
	digits := 6
	last := -1
	streaks := make([]int, 0)
	curStreak := 1
	for p := 0; p < digits; p++ {
		d := digit(pass, p)
		if d < last {
			return false
		}
		if d == last {
			curStreak++
		} else {
			streaks = append(streaks, curStreak)
			curStreak = 1
		}
		last = d
	}
	streaks = append(streaks, curStreak)
	for _, s := range streaks {
		if s == 2 {
			return true
		}
	}
	return false
}

func digit(num int64, place int) int {
	ds := string(strconv.FormatInt(int64(num), 10)[place])
	d, _ := strconv.ParseInt(ds, 10, 64)
	return int(d)
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	inputs04 := utils.ParseInputInts(input, "-")
	min := inputs04[0]
	max := inputs04[1]

	var numPassed int64
	for n := min; n <= max; n++ {
		if isValid(n) {
			numPassed++
		}
	}

	return fmt.Sprint(numPassed)
}
