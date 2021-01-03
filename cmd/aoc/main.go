package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/drtchops/aoc/solutions"
)

var USAGE = "Usage: aoc <year(2015-2020)> <day(1-25)> <part(a|b)>"

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println(USAGE)
		return
	}

	year, err := strconv.Atoi(args[0])
	if err != nil || year < 2015 || year > 2020 {
		fmt.Println(USAGE)
		return
	}

	day, err := strconv.Atoi(args[1])
	if err != nil || day < 1 || day > 25 {
		fmt.Println(USAGE)
		return
	}

	part := strings.ToLower(args[2])
	if part != "a" && part != "b" {
		fmt.Println(USAGE)
		return
	}

	fmt.Printf("Solving year %d day %d part %s\n", year, day, part)

	t := time.Now()
	answer, err := solutions.Solve(year, day, part)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Took %.2f seconds\n", time.Since(t).Seconds())
	fmt.Printf("Answer: %s\n", answer)
}
