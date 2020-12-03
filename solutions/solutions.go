package solutions

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/drtchops/aoc/solutions/day01"
	"github.com/drtchops/aoc/solutions/day02"
	"github.com/drtchops/aoc/solutions/day03"
	"github.com/drtchops/aoc/solutions/day04"
	"github.com/drtchops/aoc/solutions/day05"
	"github.com/drtchops/aoc/solutions/day06"
	"github.com/drtchops/aoc/solutions/day07"
	"github.com/drtchops/aoc/solutions/day08"
	"github.com/drtchops/aoc/solutions/day09"
	"github.com/drtchops/aoc/solutions/day10"
	"github.com/drtchops/aoc/solutions/day11"
	"github.com/drtchops/aoc/solutions/day12"
	"github.com/drtchops/aoc/solutions/day13"
	"github.com/drtchops/aoc/solutions/day14"
	"github.com/drtchops/aoc/solutions/day15"
	"github.com/drtchops/aoc/solutions/day16"
	"github.com/drtchops/aoc/solutions/day17"
	"github.com/drtchops/aoc/solutions/day18"
	"github.com/drtchops/aoc/solutions/day19"
	"github.com/drtchops/aoc/solutions/day20"
	"github.com/drtchops/aoc/solutions/day21"
	"github.com/drtchops/aoc/solutions/day22"
	"github.com/drtchops/aoc/solutions/day23"
	"github.com/drtchops/aoc/solutions/day24"
	"github.com/drtchops/aoc/solutions/day25"
)

type Solver interface {
	SolveA() string
	SolveB() string
}

func GetSolver(day int64) Solver {
	input := getInput(day)

	switch day {
	case 1:
		return day01.New(input)
	case 2:
		return day02.New(input)
	case 3:
		return day03.New(input)
	case 4:
		return day04.New(input)
	case 5:
		return day05.New(input)
	case 6:
		return day06.New(input)
	case 7:
		return day07.New(input)
	case 8:
		return day08.New(input)
	case 9:
		return day09.New(input)
	case 10:
		return day10.New(input)
	case 11:
		return day11.New(input)
	case 12:
		return day12.New(input)
	case 13:
		return day13.New(input)
	case 14:
		return day14.New(input)
	case 15:
		return day15.New(input)
	case 16:
		return day16.New(input)
	case 17:
		return day17.New(input)
	case 18:
		return day18.New(input)
	case 19:
		return day19.New(input)
	case 20:
		return day20.New(input)
	case 21:
		return day21.New(input)
	case 22:
		return day22.New(input)
	case 23:
		return day23.New(input)
	case 24:
		return day24.New(input)
	case 25:
		return day25.New(input)
	}

	return nil
}

func getInput(day int64) string {
	label := strconv.FormatInt(day, 10)
	if len(label) == 1 {
		label = "0" + label
	}

	inputBytes, err := ioutil.ReadFile(fmt.Sprintf("./solutions/day%s/input.txt", label))
	if err != nil {
		panic(err)
	}

	return string(inputBytes)
}
