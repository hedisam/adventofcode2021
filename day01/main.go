package main

import (
	"fmt"
	"math"

	"github.com/srowles/adventofcode2021"
)

func main() {
	input := day1Input()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func day1Input() []int {
	input := adventofcode2021.MustInputFromWebsite("1")
	return adventofcode2021.MustIntList(input)
}

func part1(measurements []int) int {
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

func part2(measurements []int) int {
	count := 0
	previous := math.MaxInt64
	for i := 0; i < len(measurements)-2; i++ {
		window := 0
		for j := i; j < i+3; j++ {
			window += measurements[j]
		}
		if window > previous {
			count++
		}
		previous = window
	}
	return count
}
