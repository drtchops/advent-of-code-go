package solutions

import (
	"fmt"

	"github.com/drtchops/advent-of-code-go/solutions/year2015"
	"github.com/drtchops/advent-of-code-go/solutions/year2019"
	"github.com/drtchops/advent-of-code-go/solutions/year2020"
	"github.com/drtchops/advent-of-code-go/solutions/year2021"
)

func Solve(year, day int, part string) (string, error) {
	switch year {
	case 2015:
		return year2015.Solve(day, part)
	case 2019:
		return year2019.Solve(day, part)
	case 2020:
		return year2020.Solve(day, part)
	case 2021:
		return year2021.Solve(day, part)
	default:
		return "", fmt.Errorf("cannot find solver for year %d", year)
	}
}
