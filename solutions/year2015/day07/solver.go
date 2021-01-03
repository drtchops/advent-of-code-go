package day07

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type wire struct {
	ID           string
	Resolved     bool
	Signal       uint16
	Source       []string
	Dependencies []string
}

func parse(input string) map[string]wire {
	lines := strings.Split(input, "\n")
	wires := make(map[string]wire)

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		source := parts[0]
		id := parts[1]

		sourceParts := strings.Split(source, " ")
		dependencies := make([]string, 0)
		switch len(sourceParts) {
		case 1:
			if _, err := strconv.Atoi(sourceParts[0]); err != nil {
				dependencies = append(dependencies, sourceParts[0])
			}
		case 2:
			if _, err := strconv.Atoi(sourceParts[1]); err != nil {
				dependencies = append(dependencies, sourceParts[1])
			}
		case 3:
			if _, err := strconv.Atoi(sourceParts[0]); err != nil {
				dependencies = append(dependencies, sourceParts[0])
			}
			if _, err := strconv.Atoi(sourceParts[2]); err != nil {
				dependencies = append(dependencies, sourceParts[2])
			}
		}

		wires[id] = wire{
			ID:           id,
			Source:       sourceParts,
			Dependencies: dependencies,
		}
	}

	return wires
}

func getValue(wires map[string]wire, label string) uint16 {
	if i, err := strconv.Atoi(label); err == nil {
		return uint16(i)
	}

	resolve(wires, label)
	return wires[label].Signal
}

func resolve(wires map[string]wire, id string) int64 {
	w, ok := wires[id]
	if !ok {
		return 0
	}
	if w.Resolved {
		return int64(w.Signal)
	}

	for _, dep := range w.Dependencies {
		resolve(wires, dep)
	}

	var signal uint16
	switch len(w.Source) {
	case 1:
		signal = getValue(wires, w.Source[0])
	case 2:
		if w.Source[0] == "NOT" {
			signal = ^getValue(wires, w.Source[1])
		}
	case 3:
		left := getValue(wires, w.Source[0])
		right := getValue(wires, w.Source[2])
		switch w.Source[1] {
		case "AND":
			signal = left & right
		case "OR":
			signal = left | right
		case "LSHIFT":
			signal = left << right
		case "RSHIFT":
			signal = left >> right
		}
	}

	w.Signal = signal
	w.Resolved = true
	wires[w.ID] = w

	return int64(signal)
}

func (s *Solver) SolveA(input string) string {
	wires := parse(input)
	output := resolve(wires, "a")
	return fmt.Sprint(output)
}

func (s *Solver) SolveB(input string) string {
	wires := parse(input)
	a := resolve(wires, "a")

	for _, w := range wires {
		w.Resolved = false
		w.Signal = 0
		wires[w.ID] = w
	}
	bw := wires["b"]
	bw.Source = []string{fmt.Sprint(a)}
	wires["b"] = bw

	output := resolve(wires, "a")
	return fmt.Sprint(output)
}
