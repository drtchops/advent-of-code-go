package day11

import (
	"fmt"

	"github.com/drtchops/aoc/solutions/year2019/intcode"
	"github.com/drtchops/aoc/utils"
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
	term := make(chan intcode.TermSig, 10)

	go intcode.Run(prog, 0, in, out, term)

	colors := make(map[utils.Point]bool)
	pos := utils.Point{}
	colors[pos] = true
	dir := "up"
	var minX, maxX, minY, maxY int64
	done := false
	sentInput := false

	for {
		if !sentInput {
			color, ok := colors[pos]
			if !ok || !color {
				in <- 0
			} else {
				in <- 1
			}
			sentInput = true
		}

		var newColor int64
		select {
		case newColor = <-out:
		default:
			newColor = -1
		}

		if newColor != -1 {
			sentInput = false
			colors[pos] = newColor == 1

			turn := <-out
			if turn == 0 {
				if dir == "up" {
					dir = "left"
					pos.X--
				} else if dir == "down" {
					dir = "right"
					pos.X++
				} else if dir == "left" {
					dir = "down"
					pos.Y--
				} else {
					dir = "up"
					pos.Y++
				}
			} else {
				if dir == "up" {
					dir = "right"
					pos.X++
				} else if dir == "down" {
					dir = "left"
					pos.X--
				} else if dir == "left" {
					dir = "up"
					pos.Y++
				} else {
					dir = "down"
					pos.Y--
				}
			}

			if pos.X < minX {
				minX = pos.X
			}
			if pos.X > maxX {
				maxX = pos.X
			}
			if pos.Y < minY {
				minY = pos.Y
			}
			if pos.Y > maxY {
				maxY = pos.Y
			}
		}

		select {
		case <-term:
			done = true
		default:
			done = false
		}

		if done {
			break
		}
	}

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			tile := " "
			c, ok := colors[utils.Point{X: x, Y: y}]
			if ok && c {
				tile = "X"
			}
			fmt.Print(tile)
		}
		fmt.Println("")
	}

	return ""
}
