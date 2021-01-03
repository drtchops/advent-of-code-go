package day11

import (
	"fmt"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type seatState string

var (
	FLOOR    seatState = "FLOOR"
	EMPTY    seatState = "EMPTY"
	OCCUPIED seatState = "OCCUPIED"
)

type point struct {
	X, Y int
}

func parse(input string) (map[point]seatState, int, int) {
	layout := make(map[point]seatState)
	lines := strings.Split(input, "\n")
	var maxX, maxY int
	for y, line := range lines {
		spots := strings.Split(line, "")
		for x, c := range spots {
			p := point{x, y}
			var s seatState
			switch c {
			case ".":
				s = FLOOR
			case "L":
				s = EMPTY
			case "#":
				s = OCCUPIED
			}
			layout[p] = s
			maxX = x
		}
		maxY = y
	}
	return layout, maxX, maxY
}

var directions = []point{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func (s *Solver) SolveA(input string) string {
	layout, _, _ := parse(input)

	for {
		newLayout := make(map[point]seatState)
		changeCount := 0

		for pos, state := range layout {
			if state == FLOOR {
				newLayout[pos] = FLOOR
				continue
			}

			count := 0
			for _, d := range directions {
				p2 := point{pos.X + d.X, pos.Y + d.Y}
				if layout[p2] == OCCUPIED {
					count++
				}
			}

			if state == EMPTY && count == 0 {
				changeCount++
				newLayout[pos] = OCCUPIED
			} else if state == OCCUPIED && count >= 4 {
				changeCount++
				newLayout[pos] = EMPTY
			} else {
				newLayout[pos] = state
			}
		}

		layout = newLayout

		if changeCount == 0 {
			break
		}
	}

	var count int64
	for _, s := range layout {
		if s == OCCUPIED {
			count++
		}
	}

	return fmt.Sprint(count)
}

func (s *Solver) SolveB(input string) string {
	layout, maxX, maxY := parse(input)

	for {
		newLayout := make(map[point]seatState)
		changeCount := 0

		for pos, state := range layout {
			if state == FLOOR {
				newLayout[pos] = FLOOR
				continue
			}

			count := 0
			for _, d := range directions {
				offset := 1
				for {
					newX := pos.X + (d.X * offset)
					newY := pos.Y + (d.Y * offset)
					if newX < 0 || newX > maxX || newY < 0 || newY > maxY {
						break
					}

					p2 := point{newX, newY}
					if layout[p2] == FLOOR {
						offset++
						continue
					}
					if layout[p2] == OCCUPIED {
						count++
					}
					break
				}
			}

			if state == EMPTY && count == 0 {
				changeCount++
				newLayout[pos] = OCCUPIED
			} else if state == OCCUPIED && count >= 5 {
				changeCount++
				newLayout[pos] = EMPTY
			} else {
				newLayout[pos] = state
			}
		}

		layout = newLayout

		if changeCount == 0 {
			break
		}
	}

	var count int64
	for _, s := range layout {
		if s == OCCUPIED {
			count++
		}
	}

	return fmt.Sprint(count)
}
