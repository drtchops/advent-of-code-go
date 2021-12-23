package day07

import (
	"fmt"

	"github.com/drtchops/advent-of-code-go/solutions/year2019/intcode"
	"github.com/drtchops/advent-of-code-go/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	var phaseSettings = []int64{5, 6, 7, 8, 9}
	var numProgs int64 = 5

	var best int64
	inputs07 := intcode.Parse(input)

	for _, phases := range utils.Permutations(phaseSettings) {
		fmt.Println(phases)
		term := make(chan intcode.TermSig, numProgs)
		inputs := make([]chan int64, numProgs)
		var n int64
		for n = 0; n < numProgs; n++ {
			inputs[n] = make(chan int64, 1000)
		}

		var i int64
		for i = 0; i < numProgs; i++ {
			prog := make([]int64, len(inputs07))
			copy(prog, inputs07)

			input := inputs[i]
			var output chan int64
			if i == numProgs-1 {
				output = inputs[0]
			} else {
				output = inputs[i+1]
			}
			input <- phases[i]
			if i == 0 {
				input <- 0
			}

			go intcode.Run(prog, i, input, output, term)
		}

		for n = 0; n < numProgs; n++ {
			out := <-term
			if out.Err != nil {
				fmt.Println(out.Err)
				break
			}
			if out.ID == numProgs-1 && out.Output > best {
				best = out.Output
			}
		}
	}

	return fmt.Sprint(best)
}
