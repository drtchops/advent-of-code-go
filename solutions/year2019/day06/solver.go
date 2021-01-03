package day06

import (
	"fmt"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type Planet struct {
	name     string
	parent   string
	children []string
}

type Step struct {
	name  string
	steps int64
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	planets := make(map[string]Planet)

	for _, orbit := range strings.Split(input, "\n") {
		parts := strings.Split(orbit, ")")
		parent := parts[0]
		child := parts[1]

		if p, ok := planets[parent]; ok {
			p.children = append(p.children, child)
			planets[parent] = p
		} else {
			p = Planet{
				name:     parent,
				children: []string{child},
			}
			planets[parent] = p
		}

		if p, ok := planets[child]; ok {
			p.parent = parent
			planets[child] = p
		} else {
			p = Planet{
				name:   child,
				parent: parent,
			}
			planets[child] = p
		}
	}

	visited := make(map[string]bool)
	next := make([]Step, 0)

	me := planets["YOU"]
	start := planets[me.parent]
	if start.parent != "" {
		next = append(next, Step{name: start.parent})
	}
	for _, c := range start.children {
		next = append(next, Step{name: c})
	}

	for len(next) > 0 {
		s := next[0]

		if s.name == "SAN" {
			return fmt.Sprint(s.steps)
		}
		visited[s.name] = true

		p := planets[s.name]
		if p.parent != "" && !visited[p.parent] {
			next = append(next, Step{steps: s.steps + 1, name: p.parent})
		}
		for _, c := range p.children {
			if !visited[c] {
				next = append(next, Step{steps: s.steps + 1, name: c})
			}
		}
		next = next[1:]
	}

	return ""
}
