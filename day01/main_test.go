package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPart1(t *testing.T) {
	tests := map[string]struct {
		measurements          []int
		expectedIncreaseCount int
	}{
		"ex1": {
			measurements: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			expectedIncreaseCount: 7,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			increase := part1(test.measurements)
			assert.Equal(t, test.expectedIncreaseCount, increase)
		})
	}
}

func TestCountPart2(t *testing.T) {
	tests := map[string]struct {
		measurements          []int
		expectedIncreaseCount int
	}{
		"ex1": {
			measurements: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			expectedIncreaseCount: 5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			increase := part2(test.measurements)
			assert.Equal(t, test.expectedIncreaseCount, increase)
		})
	}
}
