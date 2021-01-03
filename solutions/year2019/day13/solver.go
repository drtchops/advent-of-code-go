package day13

import (
	"fmt"
	"time"

	"github.com/drtchops/aoc/solutions/year2019/intcode"
	"github.com/drtchops/aoc/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

var WIDTH int64 = 43
var HEIGHT int64 = 23

func updateState(state []int64, output chan int64, ball, paddle *utils.Point) bool {
	var changed bool

	for {
		var done bool
		var x int64
		select {
		case x = <-output:
			changed = true
		default:
			done = true
		}

		if done {
			break
		}

		y := <-output
		tid := <-output

		if x == -1 && y == 0 {
			if tid == 0 {
				fmt.Println("old score:", state[len(state)-1])
			}
			state[len(state)-1] = tid
		} else {
			state[idx(x, y)] = tid
		}

		if tid == 3 {
			paddle.X = x
			paddle.Y = y
		} else if tid == 4 {
			ball.X = x
			ball.Y = y
		}
	}

	return changed
}

func printState(state []int64) {
	fmt.Println("Score:", state[len(state)-1])

	var x int64
	var y int64
	for y = 0; y < HEIGHT; y++ {
		for x = 0; x < WIDTH; x++ {
			tile := ""
			switch state[idx(x, y)] {
			case 0:
				tile = " "
			case 1:
				tile = "█"
			case 2:
				tile = "▒"
			case 3:
				tile = "_"
			case 4:
				tile = "o"
			}
			fmt.Print(tile)
		}
		fmt.Println("")
	}
}

func doInput(input chan int64, ball, paddle *utils.Point) bool {
	// char, _, err := keyboard.GetSingleKey()
	// if err != nil {
	// 	panic(err)
	// }

	// switch char {
	// case 'q':
	// 	return true
	// case 'a':
	// 	input <- -1
	// case 'd':
	// 	input <- 1
	// default:
	// 	input <- 0
	// }

	if ball.X > paddle.X {
		input <- 1
	} else if ball.X < paddle.X {
		input <- -1
	} else {
		input <- 0
	}

	return false
}

func idx(x, y int64) int64 {
	return (y * WIDTH) + x
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	prog := intcode.Parse(input)
	prog[0] = 2
	in := make(chan int64, 10000)
	out := make(chan int64, 10000)
	term := make(chan intcode.TermSig, 5)

	go intcode.Run(prog, 0, in, out, term)

	state := make([]int64, WIDTH*HEIGHT+1)
	var needsInput bool
	ball := utils.Point{}
	paddle := utils.Point{}

	for {
		var done bool
		changed := updateState(state, out, &ball, &paddle)

		if changed {
			needsInput = true
			// printState13(state)
		} else if needsInput {
			needsInput = false
			if quit := doInput(in, &ball, &paddle); quit {
				done = true
			}
			time.Sleep(5 * time.Millisecond)
		} else {
			select {
			case <-term:
				done = true
			default:
				done = false
			}
		}

		if done {
			break
		}
	}

	for {
		changed := updateState(state, out, &ball, &paddle)
		if !changed {
			break
		}
		printState(state)
		fmt.Println("did not drain")
	}
	return fmt.Sprint("Score:", state[len(state)-1])
}
