package day02

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type box struct {
	l, w, h int64
}

func (b *box) Area() int64 {
	return 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l
}

func (b *box) Volume() int64 {
	return b.l * b.w * b.h
}

func (b *box) SmallestSideArea() int64 {
	sides := []int64{b.l, b.w, b.h}
	sort.Slice(sides, func(i, j int) bool { return sides[i] < sides[j] })
	return sides[0] * sides[1]
}

func (b *box) SmallestSidePerimeter() int64 {
	sides := []int64{b.l, b.w, b.h}
	sort.Slice(sides, func(i, j int) bool { return sides[i] < sides[j] })
	return (2 * sides[0]) + (2 * sides[1])
}

func parse(input string) []box {
	lines := strings.Split(input, "\n")
	boxes := make([]box, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "x")
		l, _ := strconv.ParseInt(parts[0], 10, 32)
		w, _ := strconv.ParseInt(parts[1], 10, 32)
		h, _ := strconv.ParseInt(parts[2], 10, 32)
		boxes[i] = box{l, w, h}
	}
	return boxes
}

func (s *Solver) SolveA(input string) string {
	boxes := parse(input)
	var total int64
	for _, b := range boxes {
		total += b.Area() + b.SmallestSideArea()
	}
	return fmt.Sprint(total)
}

func (s *Solver) SolveB(input string) string {
	boxes := parse(input)
	var total int64
	for _, b := range boxes {
		total += b.Volume() + b.SmallestSidePerimeter()
	}
	return fmt.Sprint(total)
}
