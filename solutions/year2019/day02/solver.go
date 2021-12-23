package day02

import (
	"fmt"

	"github.com/drtchops/advent-of-code-go/solutions/year2019/intcode"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

var WANT int64 = 19690720

func seed(input []int64, noun, verb int64) []int64 {
	prog := make([]int64, len(input))
	copy(prog, input)
	prog[1] = noun
	prog[2] = verb
	return prog
}

func compute(prog []int64) (int64, error) {
	ptr := 0

	for {
		opt := prog[ptr]
		if opt == 99 {
			break
		}
		if opt != 1 && opt != 2 {
			return 0, fmt.Errorf("unknown opt %d at addr %d", opt, ptr)
		}
		if len(prog) < ptr+3 {
			return 0, fmt.Errorf("not enough values")
		}

		addr1 := prog[ptr+1]
		addr2 := prog[ptr+2]
		addr3 := prog[ptr+3]
		val1 := prog[addr1]
		val2 := prog[addr2]

		var val int64
		if opt == 1 {
			val = val1 + val2
		} else {
			val = val1 * val2
		}

		prog[addr3] = val
		ptr += 4
	}

	return prog[0], nil
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	progInput := intcode.Parse(input)

	var n int64
	var v int64

	for n = 0; n < 100; n++ {
		for v = 0; v < 100; v++ {
			prog := seed(progInput, n, v)
			output, err := compute(prog)
			if err == nil && output == WANT {
				return fmt.Sprintf("%02d%02d", n, v)
			}
		}
	}

	return ""
}
