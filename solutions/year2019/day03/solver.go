package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/drtchops/aoc/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func follow(cmd string, p utils.Point) []utils.Point {
	d := string(cmd[0])
	n, _ := strconv.ParseInt(cmd[1:], 10, 64)
	points := make([]utils.Point, n)

	var i int64
	if d == "U" {
		for i = 1; i <= n; i++ {
			points[i-1] = utils.Point{
				X: p.X,
				Y: p.Y + i,
			}
		}
	} else if d == "D" {
		for i = 1; i <= n; i++ {
			points[i-1] = utils.Point{
				X: p.X,
				Y: p.Y - i,
			}
		}
	} else if d == "R" {
		for i = 1; i <= n; i++ {
			points[i-1] = utils.Point{
				X: p.X + i,
				Y: p.Y,
			}
		}
	} else if d == "L" {
		for i = 1; i <= n; i++ {
			points[i-1] = utils.Point{
				X: p.X - i,
				Y: p.Y,
			}
		}
	}

	return points
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	inputs03 := strings.Split(input, "\n")
	wire1 := strings.Split(inputs03[0], ",")
	wire2 := strings.Split(inputs03[1], ",")

	points1 := make(map[utils.Point]int64)
	p := utils.Point{}
	var steps int64
	for _, cmd := range wire1 {
		points := follow(cmd, p)
		for i, pp := range points {
			if _, ok := points1[pp]; !ok {
				points1[pp] = steps + int64(i) + 1
			}
		}
		p = points[len(points)-1]
		steps += int64(len(points))
	}

	intersections := make(map[utils.Point]int64)
	points2 := make(map[utils.Point]int64)
	p = utils.Point{}
	steps = 0
	for _, cmd := range wire2 {
		points := follow(cmd, p)
		for i, pp := range points {
			d2 := steps + int64(i) + 1
			if _, ok := points2[pp]; !ok {
				points2[pp] = d2
			}
			if d1, ok := points1[pp]; ok {
				intersections[pp] = d1 + d2
			}
		}
		p = points[len(points)-1]
		steps += int64(len(points))
	}

	var best int64 = 99999
	for _, d := range intersections {
		if d < best {
			best = d
		}
	}

	return fmt.Sprint(best)
}
