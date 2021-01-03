package day03

import (
	"fmt"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type point struct {
	X, Y int64
}

func (p *point) Add(x, y int64) point {
	return point{p.X + x, p.Y + y}
}

func parse(input string) [][]bool {
	lines := strings.Split(input, "\n")
	trees := make([][]bool, len(lines))

	for i, line := range lines {
		trees[i] = make([]bool, len(line))
		for j, c := range line {
			if c == '#' {
				trees[i][j] = true
			}
		}
	}

	return trees
}

func testSlope(right, down int64, tree [][]bool) int64 {
	if down <= 0 {
		return 0
	}

	p := point{0, 0}
	var hit int64
	for {
		p = p.Add(right, down)
		if p.Y >= int64(len(tree)) {
			break
		}

		row := tree[p.Y]
		if row[p.X%int64(len(row))] {
			hit++
		}
	}

	return hit
}

func (s *Solver) SolveA(input string) string {
	tree := parse(input)
	hit := testSlope(3, 1, tree)
	return fmt.Sprint(hit)
}

func (s *Solver) SolveB(input string) string {
	tree := parse(input)
	var answer int64 = 1
	slopes := []point{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, slope := range slopes {
		hit := testSlope(slope.X, slope.Y, tree)
		answer *= hit
	}

	return fmt.Sprint(answer)
}
