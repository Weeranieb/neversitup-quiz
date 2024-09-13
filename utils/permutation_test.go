package utils_test

import (
	"neverquiz/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	tests := []struct {
		input  string
		output []string
	}{
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"ab", []string{"ab", "ba"}},
		{"a", []string{"a"}},
	}

	for _, test := range tests {
		if got := utils.Permutation(test.input); !assert.ElementsMatch(t, got, test.output) {
			assert.Fail(t, "Permutation(%v) = %v; want %v", test.input, got, test.output)
		}
	}
}
