package day12

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

type command struct {
	Action rune
	Value  int
}

var (
	NORTH   = 'N'
	SOUTH   = 'S'
	EAST    = 'E'
	WEST    = 'W'
	LEFT    = 'L'
	RIGHT   = 'R'
	FORWARD = 'F'
)

var leftDirections = []rune{NORTH, WEST, SOUTH, EAST}
var rightDirections = []rune{NORTH, EAST, SOUTH, WEST}

func indexOf(s []rune, e rune) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

func parse(input string) []command {
	lines := strings.Split(input, "\n")
	commands := make([]command, len(lines))

	for i, line := range lines {
		val, _ := strconv.Atoi(line[1:])
		commands[i] = command{
			Action: rune(line[0]),
			Value:  val,
		}
	}

	return commands
}

func move(p point, d rune, v int) point {
	switch d {
	case NORTH:
		p.Y += v
	case SOUTH:
		p.Y -= v
	case EAST:
		p.X += v
	case WEST:
		p.X -= v
	}
	return p
}

func turn(facing, dir rune, v int) rune {
	turns := v / 90
	directions := leftDirections
	if dir == RIGHT {
		directions = rightDirections
	}
	idx := indexOf(directions, facing) + turns
	return directions[idx%4]
}

func rotate(wp point, dir rune, v int) point {
	if dir == RIGHT {
		v = 360 - v
	}
	angle := float64(v) * (math.Pi / 180)
	x := float64(wp.X)
	y := float64(wp.Y)
	s := math.Sin(angle)
	c := math.Cos(angle)
	newX := math.Round(x*c - y*s)
	newY := math.Round(y*c + x*s)
	return point{int(newX), int(newY)}
}

func (s *Solver) SolveA(input string) string {
	commands := parse(input)

	pos := point{0, 0}
	dir := EAST
	for _, c := range commands {
		switch c.Action {
		case LEFT:
			fallthrough
		case RIGHT:
			dir = turn(dir, c.Action, c.Value)
		case FORWARD:
			pos = move(pos, dir, c.Value)
		default:
			pos = move(pos, c.Action, c.Value)
		}
	}

	dist := math.Abs(float64(pos.X)) + math.Abs(float64(pos.Y))
	return fmt.Sprint(dist)
}

func (s *Solver) SolveB(input string) string {
	commands := parse(input)

	ship := point{0, 0}
	wp := point{10, 1}

	for _, c := range commands {
		switch c.Action {
		case NORTH:
			wp.Y += c.Value
		case SOUTH:
			wp.Y -= c.Value
		case EAST:
			wp.X += c.Value
		case WEST:
			wp.X -= c.Value
		case LEFT:
			fallthrough
		case RIGHT:
			wp = rotate(wp, c.Action, c.Value)
		case FORWARD:
			ship.X += wp.X * c.Value
			ship.Y += wp.Y * c.Value
		}
	}

	dist := math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y))
	return fmt.Sprint(dist)
}
