package day10

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/drtchops/advent-of-code-go/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func parse(input string) []utils.Point {
	lines := strings.Split(input, "\n")
	rocks := make([]utils.Point, 0)

	for y, line := range lines {
		points := strings.Split(line, "")
		for x, p := range points {
			if p == "#" {
				rocks = append(rocks, utils.Point{X: int64(x), Y: int64(y)})
			}
		}
	}

	return rocks
}

func angle(p utils.Point) float64 {
	return math.Atan2(float64(p.X), float64(p.Y))
}

func magnitude(p utils.Point) int64 {
	return int64Abs(p.X) + int64Abs(p.Y)
}

func int64Abs(n int64) int64 {
	return int64(math.Abs(float64(n)))
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	rocks := parse(input)
	rock := utils.Point{X: 11, Y: 11}
	angles := make(map[utils.Point][]utils.Point)
	keys := make([]utils.Point, 0)

	for _, other := range rocks {
		if other == rock {
			continue
		}

		dx := other.X - rock.X
		dy := other.Y - rock.Y
		gcd := utils.GCD(int64Abs(dx), int64Abs(dy))
		dx /= gcd
		dy /= gcd

		a := utils.Point{X: dx, Y: dy}

		group, ok := angles[a]
		if !ok {
			group = make([]utils.Point, 0)
			keys = append(keys, a)
		}
		group = append(group, other)
		sort.Slice(group, func(i, j int) bool {
			return magnitude(group[i]) < magnitude(group[j])
		})
		angles[a] = group
	}

	sort.Slice(keys, func(i, j int) bool {
		return angle(keys[i]) > angle(keys[j])
	})

	count := 0
	last := utils.Point{}

	for count < 200 {
		for _, k := range keys {
			rocks := angles[k]
			if len(rocks) == 0 {
				continue
			}

			last = rocks[0]
			angles[k] = rocks[1:]
			count++
			if count >= 200 {
				break
			}
		}
	}

	return fmt.Sprint((last.X * 100) + last.Y)
}
