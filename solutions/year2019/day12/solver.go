package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/drtchops/aoc/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type Match struct {
	id    int
	start int64
	stop  int64
}

func parse(input string) ([]int64, []int64, []int64) {
	re := regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)
	lines := strings.Split(input, "\n")
	posX := make([]int64, len(lines))
	posY := make([]int64, len(lines))
	posZ := make([]int64, len(lines))

	for i, line := range lines {
		pos := re.FindStringSubmatch(line)[1:]
		x, _ := strconv.ParseInt(pos[0], 10, 64)
		y, _ := strconv.ParseInt(pos[1], 10, 64)
		z, _ := strconv.ParseInt(pos[2], 10, 64)
		posX[i] = x
		posY[i] = y
		posZ[i] = z
	}

	return posX, posY, posZ
}

func findLoop(id int, points []int64, perms [][]int, matches chan Match) {
	vels := make([]int64, len(points))

	history := make(map[string]int64)
	history[key(points, vels)] = 0

	var step int64
	for step = 1; ; step++ {
		for _, pair := range perms {
			i1 := pair[0]
			i2 := pair[1]
			p1 := points[i1]
			p2 := points[i2]

			if p1 > p2 {
				vels[i1]--
				vels[i2]++
			} else if p1 < p2 {
				vels[i1]++
				vels[i2]--
			}
		}

		for i := range points {
			points[i] += vels[i]
		}

		s := key(points, vels)
		if start, ok := history[s]; ok {
			matches <- Match{id, start, step}
			break
		} else {
			history[s] = step
		}
	}
}

func key(pos []int64, vel []int64) string {
	return fmt.Sprintf("%d:%d:%d:%d:%d:%d:%d:%d", pos[0], pos[1], pos[2], pos[3], vel[0], vel[1], vel[2], vel[3])
}

func hasMatch(steps [][]int64) bool {
	for _, x := range steps[0] {
		for _, y := range steps[0] {
			if x != y {
				continue
			}

			for _, z := range steps[0] {
				if x == y && y == z {
					return true
				}
			}
		}
	}

	return false
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	posX, posY, posZ := parse(input)
	matches := make(chan Match, 100)
	perms := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 2},
		{1, 3},
		{2, 3},
	}

	go findLoop(0, posX, perms, matches)
	go findLoop(1, posY, perms, matches)
	go findLoop(2, posZ, perms, matches)

	stops := make([]int64, 0)
	for {
		m := <-matches
		stops = append(stops, m.stop)
		if len(stops) == 3 {
			break
		}
	}

	return fmt.Sprint(utils.LCM(stops[0], stops[1], stops[2]))
}
