package yearbase

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day01"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day02"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day03"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day04"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day05"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day06"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day07"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day08"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day09"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day10"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day11"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day12"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day13"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day14"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day15"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day16"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day17"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day18"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day19"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day20"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day21"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day22"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day23"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day24"
	"github.com/drtchops/advent-of-code-go/solutions/yearbase/day25"
)

var YEAR = 2022

type Solver interface {
	SolveA(input string) string
	SolveB(input string) string
}

func Solve(day int, part string) (string, error) {
	input := getInput(day)

	var solver Solver

	switch day {
	case 1:
		solver = day01.New()
	case 2:
		solver = day02.New()
	case 3:
		solver = day03.New()
	case 4:
		solver = day04.New()
	case 5:
		solver = day05.New()
	case 6:
		solver = day06.New()
	case 7:
		solver = day07.New()
	case 8:
		solver = day08.New()
	case 9:
		solver = day09.New()
	case 10:
		solver = day10.New()
	case 11:
		solver = day11.New()
	case 12:
		solver = day12.New()
	case 13:
		solver = day13.New()
	case 14:
		solver = day14.New()
	case 15:
		solver = day15.New()
	case 16:
		solver = day16.New()
	case 17:
		solver = day17.New()
	case 18:
		solver = day18.New()
	case 19:
		solver = day19.New()
	case 20:
		solver = day20.New()
	case 21:
		solver = day21.New()
	case 22:
		solver = day22.New()
	case 23:
		solver = day23.New()
	case 24:
		solver = day24.New()
	case 25:
		solver = day25.New()
	default:
		return "", fmt.Errorf("cannot find solver for year %d day %d", YEAR, day)
	}

	switch part {
	case "a":
		return solver.SolveA(input), nil
	case "b":
		return solver.SolveB(input), nil
	default:
		return "", fmt.Errorf("cannot find solver for year %d day %d part %s", YEAR, day, part)
	}
}

func getInput(day int) string {
	label := strconv.Itoa(day)
	if len(label) == 1 {
		label = "0" + label
	}

	inputBytes, err := ioutil.ReadFile(fmt.Sprintf("./solutions/year%d/day%s/input.txt", YEAR, label))
	if err != nil {
		panic(err)
	}

	return string(inputBytes)
}
