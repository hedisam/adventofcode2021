package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountIncrease(t *testing.T) {
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
			increase := solution(test.measurements)
			assert.Equal(t, test.expectedIncreaseCount, increase)
		})
	}
}
