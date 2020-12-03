package utils

import (
	"strconv"
	"strings"
)

// ParseInputInts takes a string input and parses it to a list of numbers.
func ParseInputInts(input, sep string) []int64 {
	lines := strings.Split(input, sep)
	parsed := make([]int64, len(lines))
	for i, line := range lines {
		num, _ := strconv.ParseInt(line, 10, 64)
		parsed[i] = num
	}
	return parsed
}

// Permutations returns a list of all permutations of the given inputs.
func Permutations(arr []int64) [][]int64 {
	var helper func([]int64, int64)
	res := [][]int64{}

	helper = func(arr []int64, n int64) {
		if n == 1 {
			tmp := make([]int64, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			var i int64
			for i = 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, int64(len(arr)))
	return res
}

// GCD finds the Greatest Common Divisor via Euclidean algorithm.
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM finds the Least Common Multiple via GCD.
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
