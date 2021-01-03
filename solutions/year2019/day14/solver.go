package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type Resource struct {
	name   string
	amount int64
}

type Reaction struct {
	reagents []Resource
	result   Resource
}

func parse(input string) []Reaction {
	lines := strings.Split(input, "\n")
	reactions := make([]Reaction, len(lines))

	for i, line := range lines {
		formulaParts := strings.Split(line, " => ")
		reactionParts := strings.Split(formulaParts[0], ", ")
		reagents := make([]Resource, len(reactionParts))
		for j, r := range reactionParts {
			reagents[j] = parseResource(r)
		}
		reactions[i] = Reaction{reagents, parseResource(formulaParts[1])}
	}

	return reactions
}

func parseResource(input string) Resource {
	parts := strings.Split(input, " ")
	a, _ := strconv.ParseInt(parts[0], 10, 64)
	return Resource{parts[1], a}
}

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	reactions := parse(input)
	resourceTotals := make(map[string]int64)
	resourceUsed := make(map[string]int64)
	resourceProducers := make(map[string]Reaction)
	for _, r := range reactions {
		resourceTotals[r.result.name] = 0
		resourceUsed[r.result.name] = 0
		resourceProducers[r.result.name] = r
	}

	requirements := []Resource{
		{"FUEL", 1},
	}

	for {
		var done bool
		newReqs := make([]Resource, 0)
		for _, req := range requirements {
			if req.name == "ORE" {
				resourceTotals["ORE"] += req.amount
				resourceUsed["ORE"] += req.amount

				if resourceTotals["ORE"] >= 1000000000000 {
					done = true
					break
				}

				continue
			}

			if done {
				break
			}

			total := resourceTotals[req.name]
			used := resourceUsed[req.name] + req.amount
			if used < total {
				resourceUsed[req.name] = used
				continue
			}

			reaction := resourceProducers[req.name]
			for total < used {
				total += reaction.result.amount
				newReqs = append(newReqs, reaction.reagents...)
			}
			resourceTotals[req.name] = total
			resourceUsed[req.name] = used
		}

		if done {
			break
		}

		if len(newReqs) == 0 {
			newReqs = append(newReqs, Resource{"FUEL", 1})
		}

		requirements = newReqs
	}

	return fmt.Sprint(resourceTotals["FUEL"] - 1)
}
