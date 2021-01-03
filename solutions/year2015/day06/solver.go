package day06

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type point struct {
	X, Y int
}

func parsePoint(input string) *point {
	parts := strings.Split(input, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return &point{x, y}
}

type actionType string

const (
	TURN_ON  actionType = "turn on"
	TURN_OFF actionType = "turn off"
	TOGGLE   actionType = "toggle"
)

func (s *Solver) SolveA(input string) string {
	lights := make(map[point]bool)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		var action actionType
		lightRange := ""
		if line[:7] == string(TURN_ON) {
			lightRange = line[8:]
			action = TURN_ON
		} else if line[:8] == string(TURN_OFF) {
			lightRange = line[9:]
			action = TURN_OFF
		} else {
			lightRange = line[7:]
			action = TOGGLE
		}

		parts := strings.Split(lightRange, " through ")
		topLeft := parsePoint(parts[0])
		bottomRight := parsePoint(parts[1])

		for x := topLeft.X; x <= bottomRight.X; x++ {
			for y := topLeft.Y; y <= bottomRight.Y; y++ {
				p := point{x, y}
				if action == TURN_ON {
					lights[p] = true
				} else if action == TURN_OFF {
					lights[p] = false
				} else {
					lights[p] = !lights[p]
				}
			}
		}
	}

	var lit int64
	for _, v := range lights {
		if v {
			lit++
		}
	}

	return fmt.Sprint(lit)
}

func (s *Solver) SolveB(input string) string {
	lights := make(map[point]int64)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		var action actionType
		lightRange := ""
		if line[:7] == string(TURN_ON) {
			lightRange = line[8:]
			action = TURN_ON
		} else if line[:8] == string(TURN_OFF) {
			lightRange = line[9:]
			action = TURN_OFF
		} else {
			lightRange = line[7:]
			action = TOGGLE
		}

		parts := strings.Split(lightRange, " through ")
		topLeft := parsePoint(parts[0])
		bottomRight := parsePoint(parts[1])

		for x := topLeft.X; x <= bottomRight.X; x++ {
			for y := topLeft.Y; y <= bottomRight.Y; y++ {
				p := point{x, y}
				if action == TURN_ON {
					lights[p]++
				} else if action == TURN_OFF {
					lights[p] = int64(math.Max(0, float64(lights[p]-1)))
				} else {
					lights[p] += 2
				}
			}
		}
	}

	var brightness int64
	for _, v := range lights {
		brightness += v
	}

	return fmt.Sprint(brightness)
}
