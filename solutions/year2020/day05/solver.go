package day05

import (
	"fmt"
	"math"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type seat struct {
	Row, Column int64
}

func (s *seat) ID() int64 {
	return s.Row*8 + s.Column
}

func midpoint(min, max int64) int64 {
	return int64(math.Ceil(float64(max-min) / 2.0))
}

func parse(input string) []seat {
	lines := strings.Split(input, "\n")
	seats := make([]seat, len(lines))

	for i, line := range lines {
		var minR int64
		var maxR int64 = 127
		var minC int64
		var maxC int64 = 7

		for j, c := range line {
			if j < 7 {
				switch c {
				case 'B':
					minR += midpoint(minR, maxR)
				case 'F':
					maxR -= midpoint(minR, maxR)
				}
			} else {
				switch c {
				case 'R':
					minC += midpoint(minC, maxC)
				case 'L':
					maxC -= midpoint(minC, maxC)
				}
			}
		}

		seats[i] = seat{minR, minC}
	}

	return seats
}

func (s *Solver) SolveA(input string) string {
	seats := parse(input)
	var max int64
	for _, seat := range seats {
		id := seat.ID()
		if id > max {
			max = id
		}
	}
	return fmt.Sprint(max)
}

func (s *Solver) SolveB(input string) string {
	occupied := make(map[int64]bool)
	seats := parse(input)
	for _, seat := range seats {
		occupied[seat.ID()] = true
	}
	var mine, i int64
	for i = 1; i < 1023; i++ {
		if !occupied[i] && occupied[i-1] && occupied[i+1] {
			mine = i
			break
		}
	}
	return fmt.Sprint(mine)
}
