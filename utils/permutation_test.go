package utils_test

import (
	"neverquiz/utils"
	"testing"
)

func TestPermutation(t *testing.T) {
	tests := []struct {
		input  string
		output []string
	}{
		// {"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		// {"ab", []string{"ab", "ba"}},
		// {"a", []string{"a"}},
		{"aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}

	for _, test := range tests {
		result := utils.Permutation(test.input)
		if len(result) != len(test.output) {
			t.Errorf("Permutation(%s) = %v; want %v", test.input, result, test.output)
		}
		answerMap := make(map[string]struct{})
		for _, answer := range test.output {
			answerMap[answer] = struct{}{}
		}

		for _, result := range result {
			if _, ok := answerMap[result]; !ok {
				t.Errorf("Permutation(%s) = %v; want %v", test.input, result, test.output)
			}
		}
	}
}
