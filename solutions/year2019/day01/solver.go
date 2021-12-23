package day01

import (
	"fmt"

	"github.com/drtchops/advent-of-code-go/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func calcFuel(mass int64) int64 {
	fuel := mass/3 - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	modules := utils.ParseInputInts(input, "\n")

	var totalFuel int64
	for _, mass := range modules {
		lastFuel := calcFuel(mass)

		for lastFuel > 0 {
			totalFuel += lastFuel
			lastFuel = calcFuel(lastFuel)
		}
	}

	return fmt.Sprint(totalFuel)
}
