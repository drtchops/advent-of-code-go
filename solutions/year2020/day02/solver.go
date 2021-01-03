package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolveA(input string) string {
	entries := strings.Split(input, "\n")
	var numValid int64

	for _, line := range entries {
		parts := strings.Split(line, ": ")
		policy := parts[0]
		password := parts[1]
		policyParts := strings.Split(policy, " ")
		lengthRange := policyParts[0]
		character := rune(policyParts[1][0])
		rangeParts := strings.Split(lengthRange, "-")
		min, _ := strconv.ParseInt(rangeParts[0], 10, 64)
		max, _ := strconv.ParseInt(rangeParts[1], 10, 64)

		var count int64
		for _, c := range password {
			if c == character {
				count += 1
			}
		}

		if count >= min && count <= max {
			numValid += 1
		}
	}

	return fmt.Sprint(numValid)
}

func (s *Solver) SolveB(input string) string {
	entries := strings.Split(input, "\n")
	var numValid int64

	for _, line := range entries {
		parts := strings.Split(line, ": ")
		policy := parts[0]
		password := parts[1]
		policyParts := strings.Split(policy, " ")
		lengthRange := policyParts[0]
		character := policyParts[1][0]
		indexParts := strings.Split(lengthRange, "-")
		i1, _ := strconv.ParseInt(indexParts[0], 10, 64)
		i2, _ := strconv.ParseInt(indexParts[1], 10, 64)

		c1 := password[i1-1]
		c2 := password[i2-1]
		if (c1 == character || c2 == character) && (c1 != c2) {
			numValid += 1
		}
	}

	return fmt.Sprint(numValid)
}
