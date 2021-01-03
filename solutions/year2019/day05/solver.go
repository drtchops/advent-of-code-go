package day05

import (
	"fmt"

	"github.com/drtchops/aoc/solutions/year2019/intcode"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	inputs05 := intcode.Parse(input)
	progInput := make(chan int64, 5)
	progInput <- 5
	progOutput := make(chan int64, 5)
	progTerm := make(chan intcode.TermSig, 5)
	intcode.Run(inputs05, 0, progInput, progOutput, progTerm)
	term := <-progTerm
	return fmt.Sprint(term.Output)
}
