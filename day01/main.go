package main

import (
	"fmt"
	"github.com/srowles/adventofcode2021"
	"math"
)

func main() {
	fmt.Println(solution(day1Input()))
}

func day1Input() []int {
	input := adventofcode2021.MustInputFromWebsite("1")
	return adventofcode2021.MustIntList(input)
}

func solution(measurements []int) int {
	count := 0
	previous := math.MaxInt64
	for _, depth := range measurements {
		if depth > previous {
			count++
		}
		previous = depth
	}
	return count
}
