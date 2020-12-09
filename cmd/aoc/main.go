package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/drtchops/aoc/solutions"
)

var USAGE = "Usage: aoc <day(1-25)> <part(a|b)>"

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println(USAGE)
		return
	}

	day, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil || day < 1 || day > 25 {
		fmt.Println(USAGE)
		return
	}

	part := strings.ToLower(args[1])
	if part != "a" && part != "b" {
		fmt.Println(USAGE)
		return
	}

	fmt.Printf("Solving day %d part %s\n", day, part)

	t := time.Now()
	answer, err := solutions.Solve(day, part)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Took %.2f seconds\n", time.Since(t).Seconds())
	fmt.Printf("Answer: %s\n", answer)
}
