package day04

import (
	"crypto/md5"
	"fmt"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func findHash(input, prefix string) int64 {
	var i int64 = 1
	for {
		key := []byte(fmt.Sprintf("%s%d", input, i))
		hash := fmt.Sprintf("%x", md5.Sum(key))
		if hash[:len(prefix)] == prefix {
			return i
		}
		i++
	}
}

func (s *Solver) SolveA(input string) string {
	hash := findHash(input, "00000")
	return fmt.Sprint(hash)
}

func (s *Solver) SolveB(input string) string {
	hash := findHash(input, "000000")
	return fmt.Sprint(hash)
}
