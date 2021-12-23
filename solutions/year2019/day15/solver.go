package day15

import (
	"fmt"

	"github.com/drtchops/advent-of-code-go/solutions/year2019/intcode"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	prog := intcode.Parse(input)
	in := make(chan int64, 100)
	out := make(chan int64, 100)
	term := make(chan intcode.TermSig, 5)

	go intcode.Run(prog, 0, in, out, term)
	t := <-term
	return fmt.Sprint(t.Output)
}
