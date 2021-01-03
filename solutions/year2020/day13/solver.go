package day13

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

type bus struct {
	ID     int64
	Offset int64
}

func parse(input string) (int64, []int64) {
	lines := strings.Split(input, "\n")
	ts, _ := strconv.ParseInt(lines[0], 10, 64)
	buses := strings.Split(lines[1], ",")
	busIDs := make([]int64, len(buses))
	for i, b := range buses {
		id, err := strconv.ParseInt(b, 10, 64)
		if err != nil {
			busIDs[i] = -1
		} else {
			busIDs[i] = id
		}
	}
	return ts, busIDs
}

func (s *Solver) SolveA(input string) string {
	ts, busIDs := parse(input)
	var min int64
	var minID int64
	for _, id := range busIDs {
		if id == -1 {
			continue
		}

		d := int64(math.Ceil(float64(ts)/float64(id)) * float64(id))
		if min == 0 || d < min {
			min = d
			minID = id
		}
	}

	timestamp := (min - ts) * minID
	return fmt.Sprint(timestamp)
}

func (s *Solver) SolveB(input string) string {
	_, busIDs := parse(input)
	buses := make([]bus, 0)
	for i, id := range busIDs {
		if id != -1 {
			buses = append(buses, bus{id, int64(i)})
		}
	}
	sort.Slice(buses, func(i, j int) bool { return buses[i].ID > buses[j].ID })

	// Had to cheat and look this up https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Computation
	var timestamp int64
	var increment int64 = 1
	for _, b := range buses {
		for (timestamp+b.Offset)%b.ID != 0 {
			timestamp += increment
		}
		increment *= b.ID
	}

	return fmt.Sprint(timestamp)
}
