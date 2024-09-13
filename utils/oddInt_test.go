package utils_test

import (
	"neverquiz/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddInt(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{7}, 7},
		{[]int{0}, 0},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 1, 0, 1, 0}, 0},
		{[]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, 4},
		// fail case
		{[]int{0, 1, 0, 1}, 0},
	}

	for _, test := range tests {
		if got := utils.OddInt(test.input); got != test.output {
			assert.Fail(t, "OddInt(%v) = %v; want %v", test.input, got, test.output)
		}
	}
}
