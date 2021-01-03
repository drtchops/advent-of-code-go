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

func move(pos point, dir string) point {
	if dir == "^" {
		return point{pos.X, pos.Y - 1}
	}
	if dir == "v" {
		return point{pos.X, pos.Y + 1}
	}
	if dir == ">" {
		return point{pos.X + 1, pos.Y}
	}
	if dir == "<" {
		return point{pos.X - 1, pos.Y}
	}
	return pos
}

func (s *Solver) SolveA(input string) string {
	moves := strings.Split(input, "")
	pos := point{0, 0}
	visited := map[point]int64{pos: 1}

	for _, dir := range moves {
		pos = move(pos, dir)
		visited[pos] += 1
	}

	return fmt.Sprint(len(visited))
}

func (s *Solver) SolveB(input string) string {
	moves := strings.Split(input, "")
	santaPos := point{0, 0}
	roboPos := point{0, 0}
	visited := map[point]int64{santaPos: 2}
	tick := 0

	for _, dir := range moves {
		if tick%2 == 0 {
			santaPos = move(santaPos, dir)
			visited[santaPos] += 1
		} else {
			roboPos = move(roboPos, dir)
			visited[roboPos] += 1
		}
		tick++
	}

	return fmt.Sprint(len(visited))
}
